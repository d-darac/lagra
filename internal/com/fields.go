package com

import (
	"fmt"
	"reflect"
	"strings"
)

type FieldNames map[string]map[string]map[string]string

func FieldNamesFromTags(fieldNames FieldNames, key string, s any) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("[com.fields] invalid type; must be struct")
	}

	rtName := rt.Name()

	if fieldNames[rtName] == nil {
		fieldNames[rtName] = make(map[string]map[string]string)
	}
	if fieldNames[rtName][key] == nil {
		fieldNames[rtName][key] = make(map[string]string)
	}

	for f := range rt.Fields() {
		tag := strings.Split(f.Tag.Get(key), ",")[0]
		if tag == "" || tag == "-" {
			continue
		}
		fieldNames[rtName][key][tag] = f.Name
	}
}

func getByTag(fieldNames FieldNames, key, tag string, s any) (string, error) {
	rt := reflect.TypeOf(s)
	if rt == nil {
		return "", fmt.Errorf("[com.fields] s must not be nil")
	}
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		return "", fmt.Errorf("[com.fields] invalid type %T; must be a struct", s)
	}
	return fieldNames[rt.Name()][key][tag], nil
}
