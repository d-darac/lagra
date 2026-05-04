package items

import (
	"github.com/d-darac/lagra/internal/com"
)

type ListItemsParams struct {
	*com.PaginationParams
	Active        *bool          `json:"active" validate:"omitnil"`
	CreatedAt     *com.TimeRange `json:"created_at" validate:"omitnil"`
	Description   *string        `json:"description" validate:"omitnil"`
	Group         *string        `json:"group" validate:"omitnil,id"`
	HasVariants   *bool          `json:"has_variants" validate:"omitnil"`
	Inventory     *string        `json:"inventory" validate:"omitnil,id"`
	Name          *string        `json:"name" validate:"omitnil"`
	ParentItem    *string        `json:"parent_item" validate:"omitnil,id"`
	PriceAmount   *int           `json:"price_amount" validate:"omitnil"`
	PriceCurrency *string        `json:"price_currency" validate:"omitnil,currency"`
	Type          *string        `json:"type" validate:"omitnil,itemtype"`
	UpdatedAt     *com.TimeRange `json:"updated_at" validate:"omitnil"`
	Variant       *bool          `json:"variant" validate:"omitnil"`
	Expand        []string       `json:"expand" validate:"omitnil,dive,oneof=group identifiers inventory"`
}
