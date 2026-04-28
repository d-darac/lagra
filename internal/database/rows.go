package database

import (
	"database/sql"
	"time"

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
