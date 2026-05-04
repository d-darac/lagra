package items

import (
	"context"
	"errors"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/services/items"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/jackc/pgx/v5"
)

func (h Handlers) Resolve(ctx context.Context, ids []id.ID, accountID id.ID) (map[id.ID]com.Resource, error) {
	items, err := h.Services.Items.GetByIDs(ctx, items.GetByIDsParams{
		AccountID: accountID,
		ItemIDs:   ids,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	resources := make(map[id.ID]com.Resource)
	for k, v := range items {
		resources[k] = v
	}
	return resources, nil
}
