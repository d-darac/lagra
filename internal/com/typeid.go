package com

import (
	"bytes"
	"encoding/json"

	"github.com/google/uuid"
	"go.jetify.com/typeid/v2"
)

var jsonNull = []byte("null")

type NullTypeID struct {
	TypeID typeid.TypeID
	Valid  bool // Valid is true if TypeID is not NULL
}

func (tid NullTypeID) MarshalJSON() ([]byte, error) {
	if tid.Valid {
		return json.Marshal(tid.TypeID)
	}
	return jsonNull, nil
}

func (tid *NullTypeID) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*tid = NullTypeID{}
		return nil
	}
	err := json.Unmarshal(data, &tid.TypeID)
	tid.Valid = err == nil
	return err
}

func FromNullUUID(prefix string, nullUUID uuid.NullUUID) (NullTypeID, error) {
	nullTypeID := NullTypeID{
		Valid: nullUUID.Valid,
	}
	if nullTypeID.Valid {
		typeID, err := typeid.FromUUID(prefix, nullUUID.UUID.String())
		if err != nil {
			return nullTypeID, err
		}
		nullTypeID.TypeID = typeID
	}
	return nullTypeID, nil
}

func FromStringPtr(v *string) (NullTypeID, error) {
	nullTypeID := NullTypeID{
		Valid: v != nil,
	}
	if nullTypeID.Valid {
		typeid, err := typeid.Parse(*v)
		if err != nil {
			return nullTypeID, err
		}
		nullTypeID.TypeID = typeid
	}
	return nullTypeID, nil
}

func ToNullUUID(nullTypeID NullTypeID) uuid.NullUUID {
	nullUUID := uuid.NullUUID{
		Valid: nullTypeID.Valid,
	}
	if nullUUID.Valid {
		nullUUID.UUID = uuid.MustParse(nullTypeID.TypeID.UUID())
	}
	return nullUUID
}

func ToUUID(typeID typeid.TypeID) uuid.UUID {
	return uuid.MustParse(typeID.UUID())
}
