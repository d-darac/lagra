package i32

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

var jsonNull = []byte("null")

type NullInt32 sql.NullInt32

func (i NullInt32) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int32)
	}
	return jsonNull, nil
}

func (i *NullInt32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*i = NullInt32{}
		return nil
	}
	err := json.Unmarshal(data, &i.Int32)
	i.Valid = err == nil
	return err
}

func ParsePtr(v *int) NullInt32 {
	nullInt32 := NullInt32{}
	nullInt32.Valid = v != nil
	if nullInt32.Valid {
		nullInt32.Int32 = int32(*v)
	}
	return nullInt32
}
