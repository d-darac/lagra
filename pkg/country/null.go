package country

import (
	"bytes"
	"encoding/json"

	"github.com/d-darac/lagra/internal/database/sqlc"
)

var jsonNull = []byte("null")

type NullCountry sqlc.NullCountry

func (c NullCountry) MarshalJSON() ([]byte, error) {
	if c.Valid {
		return json.Marshal(c.Country)
	}
	return jsonNull, nil
}

func (c *NullCountry) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*c = NullCountry{}
		return nil
	}
	err := json.Unmarshal(data, &c.Country)
	c.Valid = err == nil
	return err
}
