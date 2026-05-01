package id

import (
	"time"

	"github.com/google/uuid"
	"go.jetify.com/typeid/v2"
)

type ID struct {
	Time   time.Time
	TypeID typeid.TypeID
}

func Parse(s string) (ID, error) {
	typeid, err := typeid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID{TypeID: typeid, Time: Time(uuid.MustParse(typeid.UUID()))}, nil
}
