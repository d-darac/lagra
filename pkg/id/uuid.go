package id

import (
	"time"

	"github.com/google/uuid"
	"go.jetify.com/typeid/v2"
)

func FromUUID(prefix, uidStr string) (ID, error) {
	typeid, err := typeid.FromUUID(prefix, uidStr)
	if err != nil {
		return ID{}, err
	}
	return ID{
		typeID: typeid,
		time:   Time(uuid.MustParse(typeid.UUID())),
		uuid:   uuid.MustParse(typeid.UUID()),
	}, nil
}

func FromNullUUID(prefix string, nullUUID uuid.NullUUID) (NullID, error) {
	nullID := NullID{Valid: nullUUID.Valid}
	if nullID.Valid {
		ID, err := FromUUID(prefix, nullUUID.UUID.String())
		if err != nil {
			return NullID{}, err
		}
		nullID.ID = ID
	}
	return nullID, nil
}

func ToNullUUID(nullID NullID) uuid.NullUUID {
	nullUUID := uuid.NullUUID{
		Valid: nullID.Valid,
	}
	if nullUUID.Valid {
		nullUUID.UUID = nullID.ID.UUID()
	}
	return nullUUID
}

func Time(uid uuid.UUID) time.Time {
	return time.Unix(uid.Time().UnixTime()).UTC()
}
