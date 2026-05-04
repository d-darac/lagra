package items

import (
	"context"
	"net/http"
	"time"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/internal/services/items"
)

type services struct {
	Items items.Service
}

type Handlers struct {
	db       database.Database
	Services services
	expCfgs  com.ExpansionConfigs
	fldNames com.FieldNames
}

func New(
	db database.Database,
	expCfgs com.ExpansionConfigs,
	fldNames com.FieldNames,
) Handlers {
	return Handlers{
		db: db,
		Services: services{
			Items: items.NewService(db),
		},
		expCfgs:  expCfgs,
		fldNames: fldNames,
	}
}

/* func (h Handlers) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())

	reqParams := CreateItemParams{}
	if err := com.JsonDecode(r, &reqParams, w); err != nil {
		com.RespondError(w, err)
		return
	}

	// if errs := h.validator.ValidateRequestParams(params); errs != nil {
	// 	com.ErrorList(w, errs)
	// 	return
	// }

	createParams, err := mapCreateParams(reqParams, accountID)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	tx, err := h.db.BeginTx(ctx)
	if err != nil {
		com.RespondError(w, err)
		return
	}
	defer tx.Rollback(ctx)

	item, err := h.Services.Items.Create(ctx, tx, createParams)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, []com.Resource{item}, reqParams.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusCreated, item)
} */

/* func (h Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())
	itemID, err := com.GetIDFromPath(r)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	tx, err := h.db.BeginTx(ctx)
	if err != nil {
		com.RespondError(w, err)
		return
	}
	defer tx.Rollback(ctx)

	h.Services.Items.Delete(ctx, tx, items.DeleteParams{AccountID: accountID, ItemID: itemID})

	if err := tx.Commit(ctx); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusOK, nil)
} */

/* func (h Handlers) Get(w http.ResponseWriter, r *http.Request) {
	accountID := com.GetAccountID(r.Context())
	itemID, err := com.GetIDFromPath(r)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	params := GetItemParams{}
	if err := com.JsonDecode(r, &params, w); err != nil {
		com.RespondError(w, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	item, err := h.Services.Items.Get(ctx, items.GetParams{AccountID: accountID, ItemID: itemID})
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, []com.Resource{item}, params.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusOK, item)
} */

func (h Handlers) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())

	reqParams := ListItemsParams{PaginationParams: &com.PaginationParams{}, CreatedAt: &com.TimeRange{}, UpdatedAt: &com.TimeRange{}}
	if err := com.JsonDecode(r, &reqParams, w); err != nil {
		com.RespondError(w, err)
		return
	}

	listParams, err := mapListParams(reqParams, accountID)
	items, hasMore, err := h.Services.Items.List(ctx, listParams)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, items, reqParams.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	res := com.NewListResponse(items, r.URL.Path, hasMore)

	com.RespondJSON(w, http.StatusOK, res)
}

/* func (h Handlers) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())
	itemID, err := com.GetIDFromPath(r)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	reqParams := UpdateItemParams{}
	if err := com.JsonDecode(r, &reqParams, w); err != nil {
		com.RespondError(w, err)
		return
	}

	// if errs := h.validator.ValidateRequestParams(params); errs != nil {
	// 	com.ErrorList(w, errs)
	// 	return
	// }

	updateParams, err := mapUpdateParams(reqParams, itemID, accountID)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	tx, err := h.db.BeginTx(ctx)
	if err != nil {
		com.RespondError(w, err)
		return
	}
	defer tx.Rollback(ctx)

	updatedItem, err := h.Services.Items.Update(ctx, tx, updateParams)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, []com.Resource{updatedItem}, reqParams.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusOK, updatedItem)
} */

func (h Handlers) expand(ctx context.Context, resources []com.Resource, fields []string) error {
	if len(fields) > 0 {
		expansions := com.ParseExpansions(fields)
		expander := com.NewResourceExpander(h.expCfgs, h.fldNames, 4)
		if err := expander.Expand(ctx, string(com.ResourceItem), resources, expansions, 0); err != nil {
			return err
		}
	}
	return nil
}
