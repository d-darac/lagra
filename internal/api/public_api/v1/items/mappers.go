package items

import (
	"github.com/d-darac/lagra/internal/services/items"
	"github.com/d-darac/lagra/pkg/boolean"
	"github.com/d-darac/lagra/pkg/currency"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	itemtype "github.com/d-darac/lagra/pkg/item_type"
	"github.com/d-darac/lagra/pkg/str"
	"github.com/d-darac/lagra/pkg/t"
)

func mapListParams(params ListItemsParams, accountID id.ID) (lp items.ListParams, err error) {
	lp.EndingBefore, err = id.ParseNull(params.EndingBefore)
	if err != nil {
		return
	}
	lp.Group, err = id.ParseNull(params.Group)
	if err != nil {
		return
	}
	lp.Inventory, err = id.ParseNull(params.Inventory)
	if err != nil {
		return
	}
	lp.ParentItem, err = id.ParseNull(params.ParentItem)
	if err != nil {
		return
	}
	lp.StartingAfter, err = id.ParseNull(params.StartingAfter)
	if err != nil {
		return
	}
	lp.AccountID = accountID
	lp.Active = boolean.ParseNull(params.Active)
	lp.CreatedAtGt = t.ParseNull(params.CreatedAt.Gt)
	lp.CreatedAtGte = t.ParseNull(params.CreatedAt.Gte)
	lp.CreatedAtLt = t.ParseNull(params.CreatedAt.Lt)
	lp.CreatedAtLte = t.ParseNull(params.CreatedAt.Lte)
	lp.Description = str.ParseNull(params.Description)
	lp.HasVariants = boolean.ParseNull(params.HasVariants)
	lp.Name = str.ParseNull(params.Name)
	lp.Limit = i32.ParseNull(params.Limit)
	lp.PriceAmount = i32.ParseNull(params.PriceAmount)
	lp.PriceCurrency = currency.ParseNull(params.PriceCurrency)
	lp.UpdatedAtGt = t.ParseNull(params.UpdatedAt.Gt)
	lp.UpdatedAtGte = t.ParseNull(params.UpdatedAt.Gte)
	lp.UpdatedAtLt = t.ParseNull(params.UpdatedAt.Lt)
	lp.UpdatedAtLte = t.ParseNull(params.UpdatedAt.Lte)
	lp.Type = itemtype.ParseNull(params.Type)
	lp.Variant = boolean.ParseNull(params.Variant)
	return
}
