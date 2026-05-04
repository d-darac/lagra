package models

import (
	"time"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/d-darac/lagra/pkg/str"
)

type Group struct {
	accountID   id.ID
	ID          id.ID          `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Description str.NullString `json:"description"`
	Name        string         `json:"name"`
	ParentGroup com.Expandable `json:"parent_group"`
}

func (i Group) AccountID() id.ID  { return i.accountID }
func (i Group) ResourceID() id.ID { return i.ID }

func (i *Group) MapGroupRow(row database.GroupRow, accountID id.ID) error {
	groupID, err := id.FromUUID(string(com.IDPrefixGroup), row.ID.String())
	if err != nil {
		return err
	}

	parentGroupID, err := id.FromNullUUID(string(com.IDPrefixGroup), row.ParentGroupID)
	if err != nil {
		return err
	}

	i.accountID = accountID
	i.ID = groupID
	i.CreatedAt = row.CreatedAt
	i.UpdatedAt = row.UpdatedAt
	i.Description = str.NullString(row.Description)
	i.Name = row.Name
	i.ParentGroup = com.Expandable{
		ID:   parentGroupID,
		Name: string(com.ResourceGroup),
	}

	return nil
}
