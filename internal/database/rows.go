package database

import (
	"database/sql"
	"time"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/google/uuid"
)

type GroupRow struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Description   sql.NullString
	Name          string
	ParentGroupID uuid.NullUUID
}

type ItemRow struct {
	ID                uuid.UUID
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Active            bool
	Description       sql.NullString
	GroupID           uuid.NullUUID
	HasVariants       bool
	ItemIdentifiersID uuid.NullUUID
	InventoryID       uuid.NullUUID
	Name              string
	ParentItemID      uuid.NullUUID
	PriceAmount       sql.NullInt32
	PriceCurrency     sqlc.NullCurrency
	Type              sqlc.ItemType
	Variant           bool
}
