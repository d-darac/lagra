package t

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"time"
)

var jsonNull = []byte("null")

type NullTime sql.NullTime

func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}
	return jsonNull, nil
}

func (t *NullTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*t = NullTime{}
		return nil
	}
	err := json.Unmarshal(data, &t.Time)
	t.Valid = err == nil
	return err
}

func ParseNull(v *time.Time) NullTime {
	nullInt32 := NullTime{}
	nullInt32.Valid = v != nil
	if nullInt32.Valid {
		nullInt32.Time = *v
	}
	return nullInt32
}
