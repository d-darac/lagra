package groups

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

func (s Service) Create(ctx context.Context, tx pgx.Tx, params CreateParams) (*models.Group, error) {
	dbParams := mapCreateGroupParams(params)

	qtx := sqlc.New(s.db.Pool()).WithTx(tx)
	row, err := qtx.CreateGroup(ctx, dbParams)
	if err != nil {
		return nil, err
	}

	group := &models.Group{}
	if err := group.MapGroupRow(database.GroupRow(row), params.AccountID); err != nil {
		return nil, err
	}

	return group, nil
}

func (s Service) Delete(ctx context.Context, tx pgx.Tx, params DeleteParams) error {
	dbParams := mapDeleteGroupParams(params)
	qtx := sqlc.New(s.db.Pool()).WithTx(tx)
	return qtx.DeleteGroup(ctx, dbParams)
}

func (s Service) Get(ctx context.Context, params GetParams) (*models.Group, error) {
	dbParams := mapGetGroupParams(params)

	row, err := s.db.Queries.GetGroup(ctx, dbParams)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &com.ResourceNotFoundErr{
				ID:       params.GroupID.String(),
				Resource: string(com.ResourceGroup),
			}
		}
		return nil, err
	}

	group := &models.Group{}
	if err := group.MapGroupRow(database.GroupRow(row), params.AccountID); err != nil {
		return nil, err
	}

	return group, nil
}

func (s Service) GetByIDs(ctx context.Context, params GetByIDsParams) (map[id.ID]*models.Group, error) {
	dbParams := mapGetGroupsByIDsParams(params)

	rows, err := s.Q.GetGroupsByIDs(ctx, dbParams)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &com.ResourceNotFoundErr{
				// ID:       params.GroupID.String(),
				Resource: string(com.ResourceGroup),
			}
		}
		return nil, err
	}

	groups := make(map[id.ID]*models.Group)
	for _, row := range rows {
		group := &models.Group{}
		if err := group.MapGroupRow(database.GroupRow(row), params.AccountID); err != nil {
			return nil, err
		}
		groups[group.ID] = group
	}

	return groups, nil
}

func (s Service) List(ctx context.Context, params ListParams) (groups []com.Resource, hasMore bool, err error) {
	if params.StartingAfter.Valid {
		_, err := s.Get(ctx, GetParams{
			AccountID: params.AccountID,
			GroupID:   params.StartingAfter.ID,
		})
		if err != nil {
			if rnfe, ok := errors.AsType[*com.ResourceNotFoundErr](err); ok {
				rnfe.Param = new("starting_after")
			}
			return groups, hasMore, err
		}
	}

	if params.EndingBefore.Valid {
		_, err := s.Get(ctx, GetParams{
			AccountID: params.AccountID,
			GroupID:   params.EndingBefore.ID,
		})
		if err != nil {
			if rnfe, ok := errors.AsType[*com.ResourceNotFoundErr](err); ok {
				rnfe.Param = new("ending_before")
			}
			return groups, hasMore, err
		}
	}

	dbParams := mapListGroupParams(params)

	rows, err := s.db.Queries.ListGroups(ctx, dbParams)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return groups, hasMore, nil
		}
		return groups, hasMore, err
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
		group := &models.Group{}
		if err := group.MapGroupRow(database.GroupRow(row), params.AccountID); err != nil {
			return groups, hasMore, err
		}
		groups = append(groups, group)
	}

	return groups, hasMore, err
}

func (s Service) Update(ctx context.Context, tx pgx.Tx, params UpdateParams) (*models.Group, error) {
	dbParams := mapUpdateGroupParams(params)

	qtx := sqlc.New(s.db.Pool()).WithTx(tx)
	row, err := qtx.UpdateGroup(ctx, dbParams)
	if err != nil {
		return nil, err
	}

	group := &models.Group{}
	if err := group.MapGroupRow(database.GroupRow(row), params.AccountID); err != nil {
		return nil, err
	}

	return group, nil
}
