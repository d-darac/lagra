package id

import (
	"time"

	"github.com/google/uuid"
	"go.jetify.com/typeid/v2"
)

func ToUUID(ID ID) uuid.UUID {
	return uuid.MustParse(ID.TypeID.UUID())
}

func FromUUID(prefix, uidStr string) (ID, error) {
	typeid, err := typeid.FromUUID(prefix, uidStr)
	if err != nil {
		return ID{}, err
	}
	return ID{TypeID: typeid, Time: Time(uuid.MustParse(typeid.UUID()))}, nil
}

func FromNullUUID(prefix string, nullUUID uuid.NullUUID) NullID {
	nullID := NullID{Valid: nullUUID.Valid}
	if nullID.Valid {
		typeID, _ := typeid.FromUUID(prefix, nullUUID.UUID.String())
		nullID.ID = ID{
			TypeID: typeID,
			Time:   Time(uuid.MustParse(typeID.UUID())),
		}
	}
	return nullID
}

func ToNullUUID(nullID NullID) uuid.NullUUID {
	nullUUID := uuid.NullUUID{
		Valid: nullID.Valid,
	}
	if nullUUID.Valid {
		nullUUID.UUID = uuid.MustParse(nullID.ID.TypeID.UUID())
	}
	return nullUUID
}

func Time(uid uuid.UUID) time.Time {
	return time.Unix(uid.Time().UnixTime())
}
