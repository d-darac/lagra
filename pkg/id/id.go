package id

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"go.jetify.com/typeid/v2"
)

type ID struct {
	time   time.Time
	typeID typeid.TypeID
	uuid   uuid.UUID
}

func Generate(prefix string) ID {
	typeID := typeid.MustGenerate(prefix)
	time := Time(uuid.MustParse(typeID.UUID()))
	return ID{time: time, typeID: typeID, uuid: uuid.MustParse(typeID.UUID())}
}

func Parse(s string) (ID, error) {
	typeid, err := typeid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID{
		typeID: typeid,
		time:   Time(uuid.MustParse(typeid.UUID())),
		uuid:   uuid.MustParse(typeid.UUID()),
	}, nil
}

func (i ID) MarshalJSON() ([]byte, error) { return json.Marshal(i.TypeID()) }
func (i *ID) UnmarshalJSON(data []byte) error {
	id, err := Parse(string(data))
	if err != nil {
		return err
	}
	*i = id
	return nil
}

func (i ID) String() string        { return i.typeID.String() }
func (i ID) Time() time.Time       { return i.time }
func (i ID) TypeID() typeid.TypeID { return i.typeID }
func (i ID) UUID() uuid.UUID       { return i.uuid }
