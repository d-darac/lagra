package database

import (
	"database/sql"
	"time"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/google/uuid"
)

type GroupRow struct {
	CreatedAt, UpdatedAt time.Time
	Name                 string
	Description          sql.NullString
	ParentGroupID        uuid.NullUUID
	ID                   uuid.UUID
}

type ItemRow struct {
	CreatedAt,
	UpdatedAt time.Time
	Type                                                  sqlc.ItemType
	Name                                                  string
	PriceCurrency                                         sqlc.NullCurrency
	Description                                           sql.NullString
	PriceAmount                                           sql.NullInt32
	GroupID, ItemIdentifiersID, InventoryID, ParentItemID uuid.NullUUID
	ID                                                    uuid.UUID
	Active, HasVariants, Variant                          bool
}
