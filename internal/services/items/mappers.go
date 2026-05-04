package items

import (
	"database/sql"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/google/uuid"
)

func mapCreateItemParams(params CreateParams) (dbParams sqlc.CreateItemParams) {
	dbParams.AccountID = params.AccountID.UUID()
	dbParams.Active = params.Active
	dbParams.CreatedAt = params.CreatedAt
	dbParams.Description = sql.NullString(params.Description)
	dbParams.GroupID = id.ToNullUUID(params.Group)
	dbParams.HasVariants = params.HasVariants
	dbParams.ID = params.ItemID.UUID()
	dbParams.InventoryID = id.ToNullUUID(params.Inventory)
	dbParams.Name = params.Name
	dbParams.PriceAmount = sql.NullInt32(params.PriceAmount)
	dbParams.PriceCurrency = sqlc.NullCurrency(params.PriceCurrency)
	dbParams.Type = params.Type
	dbParams.UpdatedAt = params.UpdatedAt
	return
}

func mapDeleteItemParams(params DeleteParams) sqlc.DeleteItemParams {
	return sqlc.DeleteItemParams{
		AccountID: params.AccountID.UUID(),
		ID:        params.ItemID.UUID(),
	}
}

func mapGetItemParams(params GetParams) sqlc.GetItemParams {
	return sqlc.GetItemParams{
		AccountID: params.AccountID.UUID(),
		ID:        params.ItemID.UUID(),
	}
}

func mapGetItemsByIDsParams(params GetByIDsParams) (dbParams sqlc.GetItemsByIDsParams) {
	dbParams.AccountID = params.AccountID.UUID()
	uuids := make([]uuid.UUID, len(params.ItemIDs))
	for i, ID := range params.ItemIDs {
		uuids[i] = ID.UUID()
	}
	dbParams.IDs = uuids
	return
}

func mapListItemParams(params ListParams) (dbParams sqlc.ListItemsParams) {
	dbParams.AccountID = params.AccountID.UUID()
	dbParams.Active = sql.NullBool(params.Active)
	dbParams.Description = sql.NullString(params.Description)
	dbParams.GroupID = id.ToNullUUID(params.Group)
	dbParams.HasVariants = sql.NullBool(params.HasVariants)
	dbParams.InventoryID = id.ToNullUUID(params.Inventory)
	dbParams.Name = sql.NullString(params.Name)
	dbParams.ParentItemID = id.ToNullUUID(params.ParentItem)
	dbParams.PriceAmount = sql.NullInt32(params.PriceAmount)
	dbParams.PriceCurrency = sqlc.NullCurrency(params.PriceCurrency)
	dbParams.Type = sqlc.NullItemType(params.Type)
	dbParams.Variant = sql.NullBool(params.Variant)
	mapPaginationParams(params, &dbParams)
	mapTimeRangeParams(params, &dbParams)
	return
}

func mapPaginationParams(params ListParams, dbParams *sqlc.ListItemsParams) {
	dbParams.EndingBefore = id.ToNullUUID(params.EndingBefore)
	dbParams.Limit = sql.NullInt32(params.Limit)
	dbParams.StartingAfter = id.ToNullUUID(params.StartingAfter)
}

func mapTimeRangeParams(params ListParams, dbParams *sqlc.ListItemsParams) {
	dbParams.CreatedAtGt = sql.NullTime(params.CreatedAtGt)
	dbParams.CreatedAtGte = sql.NullTime(params.CreatedAtGte)
	dbParams.CreatedAtLt = sql.NullTime(params.CreatedAtLt)
	dbParams.CreatedAtLte = sql.NullTime(params.CreatedAtLte)
	dbParams.UpdatedAtGt = sql.NullTime(params.UpdatedAtGt)
	dbParams.UpdatedAtGte = sql.NullTime(params.UpdatedAtGte)
	dbParams.UpdatedAtLt = sql.NullTime(params.UpdatedAtLt)
	dbParams.UpdatedAtLte = sql.NullTime(params.UpdatedAtLte)
}

func mapUpdateItemParams(params UpdateParams) (dbParams sqlc.UpdateItemParams) {
	dbParams.AccountID = params.AccountID.UUID()
	dbParams.Active = sql.NullBool(params.Active)
	dbParams.Description = sql.NullString(params.Description)
	dbParams.GroupID = id.ToNullUUID(params.Group)
	dbParams.HasVariants = sql.NullBool(params.HasVariants)
	dbParams.ID = params.ItemID.UUID()
	dbParams.InventoryID = id.ToNullUUID(params.Inventory)
	dbParams.Name = sql.NullString(params.Name)
	dbParams.PriceAmount = sql.NullInt32(params.PriceAmount)
	dbParams.PriceCurrency = sqlc.NullCurrency(params.PriceCurrency)
	dbParams.UpdatedAt = params.UpdatedAt
	return
}
