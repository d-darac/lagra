package str

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

var jsonNull = []byte("null")

type NullString sql.NullString

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return jsonNull, nil
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*s = NullString{}
		return nil
	}
	err := json.Unmarshal(data, &s.String)
	s.Valid = err == nil
	return err
}

func ParsePtr(v *string) NullString {
	nullString := NullString{}
	nullString.Valid = v != nil
	if nullString.Valid {
		nullString.String = *v
	}
	return nullString
}
