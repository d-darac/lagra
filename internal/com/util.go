package com

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/d-darac/lagra/pkg/id"
	"github.com/google/uuid"
)

func GetAccountID(ctx context.Context) id.ID {
	accountID, _ := id.FromUUID(string(IDPrefixGroup), uuid.MustParse("019de41f-fcc1-770b-8475-ec4d7951de69").String())
	return accountID
	// return ctx.Value(CtxKeyAccountID).(tid.ID)
}

func GetIDFromPath(r *http.Request) (id.ID, error) {
	pathValue := r.PathValue("id")
	ID, err := id.Parse(pathValue)
	if err != nil {
		return ID, &InvalidIDErr{
			Param: nil,
			Value: pathValue,
		}
	}
	return ID, nil
}

func GetRequestID(ctx context.Context) id.ID {
	return ctx.Value(CtxKeyRequestID).(id.ID)
}

func GetUserID(ctx context.Context) id.ID {
	return ctx.Value(CtxKeyUserID).(id.ID)
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
