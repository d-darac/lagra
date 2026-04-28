package groups

import (
	"context"
	"errors"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/services/groups"
	"github.com/jackc/pgx/v5"
	"go.jetify.com/typeid/v2"
)

func (h Handlers) Resolve(ctx context.Context, ids []typeid.TypeID, accountID typeid.TypeID) (map[typeid.TypeID]com.Resource, error) {
	groups, err := h.Services.Groups.GetByIDs(ctx, groups.GetByIDsParams{
		AccountID: accountID,
		IDs:       ids,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	resources := make(map[typeid.TypeID]com.Resource)
	for k, v := range groups {
		resources[k] = v
	}
	return resources, nil
}
