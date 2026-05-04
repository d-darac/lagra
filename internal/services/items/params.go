package items

import (
	"time"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/d-darac/lagra/pkg/boolean"
	"github.com/d-darac/lagra/pkg/currency"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	itemtype "github.com/d-darac/lagra/pkg/item_type"
	"github.com/d-darac/lagra/pkg/str"
	"github.com/d-darac/lagra/pkg/t"
)

type CreateParams struct {
	CreatedAt, UpdatedAt time.Time
	Name                 string
	Type                 sqlc.ItemType
	PriceCurrency        currency.NullCurrency
	Description          str.NullString
	ItemID, AccountID    id.ID
	Group, Inventory     id.NullID
	PriceAmount          i32.NullInt32
	HasVariants, Active  bool
}

type DeleteParams struct {
	AccountID, ItemID id.ID
}

type GetParams struct {
	AccountID, ItemID id.ID
}

type ListParams struct {
	CreatedAtGt, CreatedAtGte,
	CreatedAtLt, CreatedAtLte,
	UpdatedAtGt, UpdatedAtGte,
	UpdatedAtLt, UpdatedAtLte t.NullTime
	Name, Description str.NullString
	Type              itemtype.NullItemType
	PriceCurrency     currency.NullCurrency
	AccountID         id.ID
	ParentItem, Inventory, Group,
	StartingAfter, EndingBefore id.NullID
	Limit, PriceAmount           i32.NullInt32
	Active, HasVariants, Variant boolean.NullBool
}

type GetByIDsParams struct {
	ItemIDs   []id.ID
	AccountID id.ID
}

type UpdateParams struct {
	UpdatedAt           time.Time
	Description, Name   str.NullString
	PriceCurrency       currency.NullCurrency
	AccountID, ItemID   id.ID
	Group, Inventory    id.NullID
	PriceAmount         i32.NullInt32
	Active, HasVariants boolean.NullBool
}
