package boolean

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

var jsonNull = []byte("null")

type NullBool sql.NullBool

func (b NullBool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	}
	return jsonNull, nil
}

func (b *NullBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*b = NullBool{}
		return nil
	}
	err := json.Unmarshal(data, &b.Bool)
	b.Valid = err == nil
	return err
}

func ParseNull(v *bool) NullBool {
	nullBool := NullBool{}
	nullBool.Valid = v != nil
	if nullBool.Valid {
		nullBool.Bool = *v
	}
	return nullBool
}
