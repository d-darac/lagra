package models

import (
	"time"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/pkg/str"
	"go.jetify.com/typeid/v2"
)

type Group struct {
	accountID   typeid.TypeID
	ID          typeid.TypeID  `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Description str.NullString `json:"description"`
	Name        string         `json:"name"`
	ParentGroup com.Expandable `json:"parent_group"`
}

func (g Group) AccountID() typeid.TypeID { return g.accountID }

func (g *Group) MapGroupRow(row database.GroupRow, accountID typeid.TypeID) error {
	groupID, err := typeid.FromUUID(string(com.TypeIDPrefixGroup), row.ID.String())
	if err != nil {
		return err
	}
	parentGroupID, err := com.FromNullUUID(string(com.TypeIDPrefixGroup), row.ParentGroupID)
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
		ID:   parentGroupID,
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
