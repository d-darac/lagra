package inventories

import (
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/internal/database/sqlc"
)

type Service struct {
	db database.Database
	Q  *sqlc.Queries
}

func NewService(db database.Database) Service {
	return Service{db: db, Q: db.Queries}
}
