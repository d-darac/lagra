package com

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"go.jetify.com/typeid/v2"
)

func GetAccountID(ctx context.Context) typeid.TypeID {
	accountID, _ := typeid.FromUUID(string(TypeIDPrefixGroup), uuid.MustParse("019da11d-ea27-764f-8294-20dffab572e5").String())
	return accountID
	// return ctx.Value(CtxKeyAccountID).(typeid.TypeID)
}

func GetIDFromPath(r *http.Request) (typeid.TypeID, error) {
	pathValue := r.PathValue("id")
	typeID, err := typeid.Parse(pathValue)
	if err != nil {
		return typeid.TypeID{}, &InvalidIDErr{
			Param: nil,
			Value: pathValue,
		}
	}
	return typeID, nil
}

func GetRequestID(ctx context.Context) typeid.TypeID {
	return ctx.Value(CtxKeyRequestID).(typeid.TypeID)
}

func GetUserID(ctx context.Context) typeid.TypeID {
	return ctx.Value(CtxKeyUserID).(typeid.TypeID)
}

func JsonDecode(r *http.Request, v any, w http.ResponseWriter) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if r.ContentLength > 0 {
		if err := decoder.Decode(v); err != nil {
			if ute, ok := errors.AsType[*json.UnmarshalTypeError](err); ok {
				jsonFieldType := ute.Type.Kind().String()
				if jsonFieldType == "slice" {
					jsonFieldType = "array"
				}
				if jsonFieldType == "struct" {
					jsonFieldType = "object"
				}
				return &JsonTypeErr{
					Param: ute.Field,
					Type:  jsonFieldType,
				}
			}
			return &JsonDecodeErr{Err: err}
		}
	}
	return nil
}
