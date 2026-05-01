package com

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"go.jetify.com/typeid/v2"
)

type Resolver interface {
	Resolve(ctx context.Context, ids []typeid.TypeID, accountID typeid.TypeID) (map[typeid.TypeID]Resource, error)
}

type Expandable struct {
	Resource Resource
	Name     string
	ID       NullTypeID
}

type ExpansionConfigs map[string]map[string]ExpansionConfig

type ExpansionConfig struct {
	// Function to batch-fetch related resources by their IDs
	// Returns a Map for O(1) lookups when merging results
	Resolver Resolver
	// Resolver func(ctx context.Context, ids []typeid.TypeID) map[typeid.TypeID]any
	// Whether this field contains an array of IDs (vs single ID)
	IsArray bool
}

type ResourceExpander struct {
	config     ExpansionConfigs
	fieldNames FieldNames
	maxDepth   int
}

type ExpansionNode struct {
	Children map[string]ExpansionNode
	Field    string
}

func (e Expandable) MarshalJSON() ([]byte, error) {
	if e.Resource != nil {
		return json.Marshal(e.Resource)
	}
	return e.ID.MarshalJSON()
}

func ParseExpansions(fields []string) map[string]ExpansionNode {
	root := make(map[string]ExpansionNode)

	if len(fields) == 0 {
		return root
	}

	for _, fld := range fields {
		parts := strings.Split(fld, ".")
		currentLvl := root

		for _, part := range parts {
			if _, ok := currentLvl[part]; !ok {
				currentLvl[part] = ExpansionNode{
					Field:    part,
					Children: make(map[string]ExpansionNode),
				}
			}
			currentLvl = currentLvl[part].Children
		}
	}

	return root
}

func NewResourceExpander(
	cfg ExpansionConfigs,
	fieldNames FieldNames,
	maxDepth int,
) ResourceExpander {
	return ResourceExpander{config: cfg, maxDepth: maxDepth, fieldNames: fieldNames}
}

func (re ResourceExpander) Expand(
	ctx context.Context,
	configName string,
	resources []Resource,
	expansions map[string]ExpansionNode,
	currentDepth int,
) error {
	if len(resources) == 0 {
		return nil
	}

	for fieldName, node := range expansions {
		currentFld := fieldName
		if currentDepth >= re.maxDepth {
			return &FieldExpandDepthErr{Field: currentFld}
		}

		fieldCfg, ok := re.config[configName][fieldName]
		if !ok {
			return &FieldUnexpandableErr{Field: fieldName}
		}

		fld, err := getByTag(re.fieldNames, "json", fieldName, resources[0])
		if err != nil {
			return err
		}

		err = re.expandField(ctx, resources, fld, fieldCfg, node, currentDepth)
		if err != nil {
			if fede, ok := errors.AsType[*FieldExpandDepthErr](err); ok {
				return &FieldExpandDepthErr{Field: currentFld + "." + fede.Field}
			}
			if fue, ok := errors.AsType[*FieldUnexpandableErr](err); ok {
				return &FieldUnexpandableErr{Field: currentFld + "." + fue.Field}
			}
			return err
		}
	}

	return nil
}

func (re ResourceExpander) expandField(
	ctx context.Context,
	resources []Resource,
	fieldName string,
	fieldCfg ExpansionConfig,
	node ExpansionNode,
	currentDepth int,
) error {
	// Collect all IDs that need to be fetched
	allIDs := make(map[typeid.TypeID]struct{})
	var accountID typeid.TypeID

	for _, resource := range resources {
		fmt.Println(currentDepth)
		accountID = resource.AccountID()
		rType := reflect.TypeOf(resource)
		if rType.Kind() != reflect.Pointer {
			continue
		}

		rPtr := reflect.ValueOf(resource)
		rVal := rPtr.Elem()

		if rVal.Kind() != reflect.Struct {
			continue
		}

		fld := rVal.FieldByName(fieldName)
		if fld.Kind() == reflect.Slice && fieldCfg.IsArray {
			for _, v := range fld.Seq2() {
				intf := v.Interface()
				if exp, ok := intf.(Expandable); ok && exp.ID.Valid {
					allIDs[exp.ID.TypeID] = struct{}{}
				}
			}
			continue
		}

		intf := fld.Interface()
		if exp, ok := intf.(Expandable); ok && exp.ID.Valid {
			allIDs[exp.ID.TypeID] = struct{}{}
		}
	}

	if len(allIDs) == 0 {
		return nil
	}

	allIDsSlice := make([]typeid.TypeID, 0)
	for k := range allIDs {
		allIDsSlice = append(allIDsSlice, k)
	}

	// Batch fetch all related resources in one query
	relatedData, err := fieldCfg.Resolver.Resolve(ctx, allIDsSlice, accountID)
	if err != nil {
		return err
	}

	// Merge fetched data back into each resource
	for _, resource := range resources {
		rType := reflect.TypeOf(resource)
		if rType.Kind() != reflect.Pointer {
			continue
		}

		rPtr := reflect.ValueOf(resource)
		rVal := rPtr.Elem()

		if rVal.Kind() != reflect.Struct {
			continue
		}

		fld := rVal.FieldByName(fieldName)
		if fld.Kind() == reflect.Slice && fieldCfg.IsArray {
			newFld := make([]Expandable, fld.Len())
			for i, v := range fld.Seq2() {
				intfI := i.Interface()
				intfV := v.Interface()
				data := relatedData[intfV.(Expandable).ID.TypeID]
				exp := Expandable{
					ID:       intfV.(Expandable).ID,
					Resource: data,
					Name:     intfV.(Expandable).Name,
				}
				newFld[intfI.(int)] = exp
			}
			rVal.FieldByName(fieldName).Set(reflect.ValueOf(newFld))
			continue
		}

		intf := fld.Interface()
		data := relatedData[intf.(Expandable).ID.TypeID]
		exp := Expandable{
			ID:       intf.(Expandable).ID,
			Resource: data,
			Name:     intf.(Expandable).Name,
		}
		rVal.FieldByName(fieldName).Set(reflect.ValueOf(exp))
	}

	if len(node.Children) > 0 {
		expandedResources := make([]Resource, 0)
		var resourceName string
		for _, resource := range resources {
			rPtr := reflect.ValueOf(resource)
			rVal := rPtr.Elem()
			if rVal.Kind() == reflect.Struct {
				fld := rVal.FieldByName(fieldName)
				if expSlc, ok := fld.Interface().([]Expandable); ok {
					resourceName = expSlc[0].Name
					for _, exp := range expSlc {
						expandedResources = append(expandedResources, exp.Resource)
					}
				}
				if exp, ok := fld.Interface().(Expandable); ok {
					expandedResources = append(expandedResources, exp.Resource)
					resourceName = exp.Name
				}
			}
		}

		if err := re.Expand(ctx, resourceName, expandedResources, node.Children, currentDepth+1); err != nil {
			return err
		}
	}
	return nil
}
