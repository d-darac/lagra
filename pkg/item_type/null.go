package itemtype

import (
	"bytes"
	"encoding/json"

	"github.com/d-darac/lagra/internal/database/sqlc"
)

var jsonNull = []byte("null")

type NullItemType sqlc.NullItemType

func (i NullItemType) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.ItemType)
	}
	return jsonNull, nil
}

func (i *NullItemType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		*i = NullItemType{}
		return nil
	}
	err := json.Unmarshal(data, &i.ItemType)
	i.Valid = err == nil
	return err
}

func ParseNull(v *string) NullItemType {
	nullItemType := NullItemType{}
	nullItemType.Valid = v != nil
	if nullItemType.Valid {
		nullItemType.ItemType = sqlc.ItemType(*v)
	}
	return nullItemType
}
