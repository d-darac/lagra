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

func (g Group) AccountID() id.ID { return g.accountID }

func (g *Group) MapGroupRow(row database.GroupRow, accountID id.ID) error {
	groupID, err := id.FromUUID(string(com.TypeIDPrefixGroup), row.ID.String())
	if err != nil {
		return err
	}
	g.accountID = accountID
	g.ID = groupID
	g.CreatedAt = row.CreatedAt
	g.UpdatedAt = row.UpdatedAt
	g.Description = str.NullString(row.Description)
	g.Name = row.Name
	g.ParentGroup = com.Expandable{
		ID:   id.FromNullUUID(string(com.TypeIDPrefixGroup), row.ParentGroupID),
		Name: string(com.ResourceGroup),
	}
	return nil
}

// func (g *Group) MapCreateGroupRow(row GroupRow) error {
// 	groupID, err := typeid.FromUUID(string(com.TypeIDPrefixGroup), row.ID.String())
// 	if err != nil {
// 		return err
// 	}
// 	parentGroupID, err := com.FromNullUUID(string(com.TypeIDPrefixGroup), row.ParentGroupID)
// 	if err != nil {
// 		return err
// 	}
// 	g.ID = groupID
// 	g.CreatedAt = row.CreatedAt
// 	g.UpdatedAt = row.UpdatedAt
// 	g.Description = str.NullString(row.Description)
// 	g.Name = row.Name
// 	g.ParentGroup = parentGroupID
// 	return nil
// }

// func (g *Group) MapUpdateGroupRow(row GroupRow) error {
// 	groupID, err := typeid.FromUUID(string(com.TypeIDPrefixGroup), row.ID.String())
// 	if err != nil {
// 		return err
// 	}
// 	parentGroupID, err := com.FromNullUUID(string(com.TypeIDPrefixGroup), row.ParentGroupID)
// 	if err != nil {
// 		return err
// 	}
// 	g.ID = groupID
// 	g.CreatedAt = row.CreatedAt
// 	g.UpdatedAt = row.UpdatedAt
// 	g.Description = str.NullString(row.Description)
// 	g.Name = row.Name
// 	g.ParentGroup = parentGroupID
// 	return nil
// }
