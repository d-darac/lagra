package items

import (
	"context"
	"errors"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/d-darac/lagra/internal/models"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	db database.Database
	Q  *sqlc.Queries
}

func NewService(db database.Database) Service {
	return Service{db: db, Q: db.Queries}
}

func (s Service) Create(ctx context.Context, tx pgx.Tx, params CreateParams) (*models.Item, error) {
	dbParams := mapCreateItemParams(params)

	qtx := sqlc.New(s.db.Pool()).WithTx(tx)
	row, err := qtx.CreateItem(ctx, dbParams)
	if err != nil {
		return nil, err
	}

	item := &models.Item{}
	if err := item.MapItemRow(
		database.ItemRow{
			ID:            row.ID,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
			Active:        row.Active,
			Description:   row.Description,
			GroupID:       row.GroupID,
			HasVariants:   row.HasVariants,
			InventoryID:   row.InventoryID,
			Name:          row.Name,
			PriceAmount:   row.PriceAmount,
			PriceCurrency: row.PriceCurrency,
			Type:          row.Type,
			Variant:       row.Variant,
		},
		params.AccountID,
	); err != nil {
		return nil, err
	}

	return item, nil
}

func (s Service) Delete(ctx context.Context, tx pgx.Tx, params DeleteParams) error {
	dbParams := mapDeleteItemParams(params)
	qtx := sqlc.New(s.db.Pool()).WithTx(tx)
	return qtx.DeleteItem(ctx, dbParams)
}

func (s Service) Get(ctx context.Context, params GetParams) (*models.Item, error) {
	dbParams := mapGetItemParams(params)

	row, err := s.db.Queries.GetItem(ctx, dbParams)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &com.ResourceNotFoundErr{
				ID:       params.ItemID.String(),
				Resource: string(com.ResourceItem),
			}
		}
		return nil, err
	}

	item := &models.Item{}
	if err := item.MapItemRow(database.ItemRow(row), params.AccountID); err != nil {
		return nil, err
	}

	return item, nil
}

func (s Service) GetByIDs(ctx context.Context, params GetByIDsParams) (map[id.ID]*models.Item, error) {
	dbParams := mapGetItemsByIDsParams(params)

	rows, err := s.Q.GetItemsByIDs(ctx, dbParams)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &com.ResourceNotFoundErr{
				// ID:       params.ItemID.String(),
				Resource: string(com.ResourceItem),
			}
		}
		return nil, err
	}

	items := make(map[id.ID]*models.Item)
	for _, row := range rows {
		item := &models.Item{}
		if err := item.MapItemRow(database.ItemRow(row), params.AccountID); err != nil {
			return nil, err
		}
		items[item.ID] = item
	}

	return items, nil
}

func (s Service) List(ctx context.Context, params ListParams) (items []com.Resource, hasMore bool, err error) {
	if params.StartingAfter.Valid {
		_, err := s.Get(ctx, GetParams{
			AccountID: params.AccountID,
			ItemID:    params.StartingAfter.ID,
		})
		if err != nil {
			if rnfe, ok := errors.AsType[*com.ResourceNotFoundErr](err); ok {
				rnfe.Param = new("starting_after")
			}
			return items, hasMore, err
		}
	}

	if params.EndingBefore.Valid {
		_, err := s.Get(ctx, GetParams{
			AccountID: params.AccountID,
			ItemID:    params.EndingBefore.ID,
		})
		if err != nil {
			if rnfe, ok := errors.AsType[*com.ResourceNotFoundErr](err); ok {
				rnfe.Param = new("ending_before")
			}
			return items, hasMore, err
		}
	}

	dbParams := mapListItemParams(params)

	rows, err := s.db.Queries.ListItems(ctx, dbParams)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return items, hasMore, nil
		}
		return items, hasMore, err
	}

	if dbParams.Limit.Valid {
		hasMore = len(rows) > int(dbParams.Limit.Int32)
	} else {
		hasMore = len(rows) > 10
	}

	if hasMore {
		if dbParams.EndingBefore.Valid {
			rows = rows[1:]
		} else {
			rows = rows[:len(rows)-1]
		}
	}

	for _, row := range rows {
		item := &models.Item{}
		if err := item.MapItemRow(database.ItemRow(row), params.AccountID); err != nil {
			return items, hasMore, err
		}
		items = append(items, item)
	}

	return items, hasMore, err
}

func (s Service) Update(ctx context.Context, tx pgx.Tx, params UpdateParams) (*models.Item, error) {
	dbParams := mapUpdateItemParams(params)

	qtx := sqlc.New(s.db.Pool()).WithTx(tx)
	row, err := qtx.UpdateItem(ctx, dbParams)
	if err != nil {
		return nil, err
	}

	item := &models.Item{}
	if err := item.MapItemRow(
		database.ItemRow{
			ID:                row.ID,
			CreatedAt:         row.CreatedAt,
			UpdatedAt:         row.UpdatedAt,
			Active:            row.Active,
			Description:       row.Description,
			GroupID:           row.GroupID,
			HasVariants:       row.HasVariants,
			ItemIdentifiersID: row.ItemIdentifiersID,
			InventoryID:       row.InventoryID,
			Name:              row.Name,
			ParentItemID:      row.ParentItemID,
			PriceAmount:       row.PriceAmount,
			PriceCurrency:     row.PriceCurrency,
			Type:              row.Type,
			Variant:           row.Variant,
		},
		params.AccountID,
	); err != nil {
		return nil, err
	}

	return item, nil
}
