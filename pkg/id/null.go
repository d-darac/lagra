package id

import (
	"bytes"
	"encoding/json"
)

var jsonNull = []byte("null")

type NullID struct {
	ID    ID
	Valid bool
}

func (id NullID) MarshalJSON() ([]byte, error) {
	if id.Valid {
		return json.Marshal(id.ID.TypeID())
	}
	return jsonNull, nil
}

func (id *NullID) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*id = NullID{}
		return nil
	}
	err := json.Unmarshal(data, &id.ID.typeID)
	id.Valid = err == nil
	return err
}

func ParseNull(s *string) (NullID, error) {
	nullID := NullID{
		Valid: s != nil,
	}
	if nullID.Valid {
		ID, err := Parse(*s)
		if err != nil {
			return nullID, err
		}
		nullID.ID = ID
	}
	return nullID, nil
}

// type NullTypeID struct {
// 	TypeID typeid.TypeID
// 	Valid  bool // Valid is true if TypeID is not NULL
// }

// func (tid NullTypeID) MarshalJSON() ([]byte, error) {
// 	if tid.Valid {
// 		return json.Marshal(tid.TypeID)
// 	}
// 	return jsonNull, nil
// }

// func (tid *NullTypeID) UnmarshalJSON(data []byte) error {
// 	if bytes.Equal(data, jsonNull) {
// 		*tid = NullTypeID{}
// 		return nil
// 	}
// 	err := json.Unmarshal(data, &tid.TypeID)
// 	tid.Valid = err == nil
// 	return err
// }

// func ParseNull(v *string) (NullTypeID, error) {
// 	nullTypeID := NullTypeID{
// 		Valid: v != nil,
// 	}
// 	if nullTypeID.Valid {
// 		typeid, err := typeid.Parse(*v)
// 		if err != nil {
// 			return nullTypeID, err
// 		}
// 		nullTypeID.TypeID = typeid
// 	}
// 	return nullTypeID, nil
// }
