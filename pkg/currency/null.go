package currency

import (
	"bytes"
	"encoding/json"

	"github.com/d-darac/lagra/internal/database/sqlc"
)

var jsonNull = []byte("null")

type NullCurrency sqlc.NullCurrency

func (c NullCurrency) MarshalJSON() ([]byte, error) {
	if c.Valid {
		return json.Marshal(c.Currency)
	}
	return jsonNull, nil
}

func (c *NullCurrency) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*c = NullCurrency{}
		return nil
	}
	err := json.Unmarshal(data, &c.Currency)
	c.Valid = err == nil
	return err
}

func ParseNull(v *string) NullCurrency {
	nullCurrency := NullCurrency{}
	nullCurrency.Valid = v != nil
	if nullCurrency.Valid {
		nullCurrency.Currency = sqlc.Currency(*v)
	}
	return nullCurrency
}
