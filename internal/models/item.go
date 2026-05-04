package models

import (
	"time"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/pkg/currency"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	itemtype "github.com/d-darac/lagra/pkg/item_type"
	"github.com/d-darac/lagra/pkg/str"
)

type ItemVariantAttributeOption struct {
	accountID id.ID
	ID        id.ID     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type ItemVariantAttribute struct {
	accountID id.ID
	ID        id.ID                         `json:"id"`
	CreatedAt time.Time                     `json:"created_at"`
	UpdatedAt time.Time                     `json:"updated_at"`
	Name      string                        `json:"name"`
	Options   []*ItemVariantAttributeOption `json:"options"`
}

type Item struct {
	accountID id.ID
	ID        id.ID     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
	// Attributes        *map[string][]string  `json:"attributes,omitempty"` // defined only if Variant=true
	Description       str.NullString        `json:"description"`
	Group             com.Expandable        `json:"group"`
	Identifiers       com.Expandable        `json:"identifiers"`
	Inventory         com.Expandable        `json:"inventory"`
	Name              string                `json:"name"`
	ParentItem        com.Expandable        `json:"parent_item"` // defined only if Variant=true
	PriceAmount       i32.NullInt32         `json:"price_amount"`
	PriceCurrency     currency.NullCurrency `json:"price_currency"`
	Variant           bool                  `json:"variant"`
	VariantAttributes *com.ListResponse     `json:"variant_attributes,omitempty"`
	Variants          *com.ListResponse     `json:"variants,omitempty"`
	Type              itemtype.ItemType     `json:"type"`
	HasVariants       bool                  `json:"-"`
}

func (g Item) AccountID() id.ID  { return g.accountID }
func (g Item) ResourceID() id.ID { return g.ID }

func (i *Item) MapItemRow(row database.ItemRow, accountID id.ID) error {
	itemID, err := id.FromUUID(string(com.IDPrefixItem), row.ID.String())
	if err != nil {
		return err
	}

	groupID, err := id.FromNullUUID(string(com.IDPrefixGroup), row.GroupID)
	if err != nil {
		return err
	}

	identifiersID, err := id.FromNullUUID(string(com.IDPrefixItemIdentifier), row.ItemIdentifiersID)
	if err != nil {
		return err
	}

	inventoryID, err := id.FromNullUUID(string(com.IDPrefixInventory), row.InventoryID)
	if err != nil {
		return err
	}

	if row.Variant {
		parentItemID, err := id.FromNullUUID(string(com.IDPrefixItem), row.ParentItemID)
		if err != nil {
			return err
		}
		i.ParentItem = com.Expandable{
			ID:   parentItemID,
			Name: string(com.ResourceItem),
		}
		// i.Attributes
	}

	i.accountID = accountID
	i.Active = row.Active
	i.CreatedAt = row.CreatedAt
	i.Description = str.NullString(row.Description)
	i.Group = com.Expandable{
		ID:   groupID,
		Name: string(com.ResourceGroup),
	}
	i.HasVariants = row.HasVariants
	i.ID = itemID
	i.Identifiers = com.Expandable{
		ID:   identifiersID,
		Name: string(com.ResourceItemIdentifier),
	}
	i.Inventory = com.Expandable{
		ID:   inventoryID,
		Name: string(com.ResourceInventory),
	}
	i.Name = row.Name
	i.PriceAmount = i32.NullInt32(row.PriceAmount)
	i.PriceCurrency = currency.NullCurrency(row.PriceCurrency)
	i.Type = itemtype.ItemType(row.Type)
	i.UpdatedAt = row.UpdatedAt
	i.Variant = row.Variant
	return nil
}
