package groups

import (
	"context"
	"errors"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/services/groups"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/jackc/pgx/v5"
)

func (h Handlers) Resolve(ctx context.Context, ids []id.ID, accountID id.ID) (map[id.ID]com.Resource, error) {
	groups, err := h.Services.Groups.GetByIDs(ctx, groups.GetByIDsParams{
		AccountID: accountID,
		GroupIDs:  ids,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	resources := make(map[id.ID]com.Resource)
	for k, v := range groups {
		resources[k] = v
	}
	return resources, nil
}
