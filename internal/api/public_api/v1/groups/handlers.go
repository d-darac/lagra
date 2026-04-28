package groups

import (
	"context"
	"net/http"
	"time"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/internal/services/groups"
)

type services struct {
	Groups groups.Service
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
			Groups: groups.NewService(db),
		},
		expCfgs:  expCfgs,
		fldNames: fldNames,
	}
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())

	reqParams := CreateGroupParams{}
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

	group, err := h.Services.Groups.Create(ctx, tx, createParams)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, []com.Resource{group}, reqParams.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusCreated, group)
}

func (h Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())
	groupID, err := com.GetIDFromPath(r)
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

	h.Services.Groups.Delete(ctx, tx, groups.DeleteParams{AccountID: accountID, GroupID: groupID})

	if err := tx.Commit(ctx); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusOK, nil)
}

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) {
	accountID := com.GetAccountID(r.Context())
	groupID, err := com.GetIDFromPath(r)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	params := GetGroupParams{}
	if err := com.JsonDecode(r, &params, w); err != nil {
		com.RespondError(w, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	group, err := h.Services.Groups.Get(ctx, groups.GetParams{AccountID: accountID, GroupID: groupID})
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, []com.Resource{group}, params.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusOK, group)
}

func (h Handlers) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())

	reqParams := ListGroupsParams{}
	if err := com.JsonDecode(r, &reqParams, w); err != nil {
		com.RespondError(w, err)
		return
	}

	listParams, err := mapListParams(reqParams, accountID)
	groups, _, err := h.Services.Groups.List(ctx, listParams)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, groups, reqParams.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusOK, groups)
}

func (h Handlers) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()

	accountID := com.GetAccountID(r.Context())
	groupID, err := com.GetIDFromPath(r)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	reqParams := UpdateGroupParams{}
	if err := com.JsonDecode(r, &reqParams, w); err != nil {
		com.RespondError(w, err)
		return
	}

	// if errs := h.validator.ValidateRequestParams(params); errs != nil {
	// 	com.ErrorList(w, errs)
	// 	return
	// }

	updateParams, err := mapUpdateParams(reqParams, groupID, accountID)
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

	updatedGroup, err := h.Services.Groups.Update(ctx, tx, updateParams)
	if err != nil {
		com.RespondError(w, err)
		return
	}

	if err := h.expand(ctx, []com.Resource{updatedGroup}, reqParams.Expand); err != nil {
		com.RespondError(w, err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		com.RespondError(w, err)
		return
	}

	com.RespondJSON(w, http.StatusOK, updatedGroup)
}

func (h Handlers) expand(ctx context.Context, resources []com.Resource, fields []string) error {
	if len(fields) > 0 {
		expansions := com.ParseExpansions(fields)
		expander := com.NewResourceExpander(h.expCfgs, h.fldNames, 4)
		if err := expander.Expand(ctx, string(com.ResourceGroup), resources, expansions, 0); err != nil {
			return err
		}
	}
	return nil
}
