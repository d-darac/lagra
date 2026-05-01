package groups

import (
	"database/sql"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/google/uuid"
)

func mapCreateGroupParams(params CreateParams) (cgp sqlc.CreateGroupParams) {
	cgp.AccountID = uuid.MustParse(params.AccountID.TypeID.UUID())
	cgp.Description = sql.NullString(params.Description)
	cgp.Name = params.Name
	cgp.ParentGroupID = id.ToNullUUID(params.ParentGroup)
	return
}

func mapDeleteGroupParams(params DeleteParams) sqlc.DeleteGroupParams {
	return sqlc.DeleteGroupParams{
		AccountID: uuid.MustParse(params.AccountID.TypeID.UUID()),
		ID:        uuid.MustParse(params.GroupID.TypeID.UUID()),
	}
}

func mapGetGroupParams(params GetParams) sqlc.GetGroupParams {
	return sqlc.GetGroupParams{
		AccountID: uuid.MustParse(params.AccountID.TypeID.UUID()),
		ID:        uuid.MustParse(params.GroupID.TypeID.UUID()),
	}
}

func mapGetGroupsByIDsParams(params GetByIDsParams) (ggbip sqlc.GetGroupsByIDsParams) {
	ggbip.AccountID = uuid.MustParse(params.AccountID.TypeID.UUID())
	uuids := make([]uuid.UUID, len(params.IDs))
	for i, tid := range params.IDs {
		uuids[i] = uuid.MustParse(tid.TypeID.UUID())
	}
	ggbip.IDs = uuids
	return
}

func mapListGroupParams(params ListParams) (lgp sqlc.ListGroupsParams) {
	lgp.AccountID = uuid.MustParse(params.AccountID.TypeID.UUID())
	lgp.Description = sql.NullString(params.Description)
	lgp.Name = sql.NullString(params.Name)
	lgp.ParentGroupID = id.ToNullUUID(params.ParentGroup)
	mapPaginationParams(params, &lgp)
	mapTimeRangeParams(params, &lgp)
	return
}

func mapPaginationParams(params ListParams, lgp *sqlc.ListGroupsParams) {
	lgp.EndingBefore = id.ToNullUUID(params.EndingBefore)
	lgp.EndingBeforeDate = sql.NullTime{Time: params.EndingBefore.ID.Time, Valid: params.EndingBefore.Valid}
	lgp.Limit = sql.NullInt32(params.Limit)
	lgp.StartingAfter = id.ToNullUUID(params.StartingAfter)
	lgp.StartingAfterDate = sql.NullTime{Time: params.StartingAfter.ID.Time, Valid: params.StartingAfter.Valid}
}

func mapTimeRangeParams(params ListParams, lgp *sqlc.ListGroupsParams) {
	lgp.CreatedAtGt = sql.NullTime(params.CreatedAtGt)
	lgp.CreatedAtGte = sql.NullTime(params.CreatedAtGte)
	lgp.CreatedAtLt = sql.NullTime(params.CreatedAtLt)
	lgp.CreatedAtLte = sql.NullTime(params.CreatedAtLte)
	lgp.UpdatedAtGt = sql.NullTime(params.UpdatedAtGt)
	lgp.UpdatedAtGte = sql.NullTime(params.UpdatedAtGte)
	lgp.UpdatedAtLt = sql.NullTime(params.UpdatedAtLt)
	lgp.UpdatedAtLte = sql.NullTime(params.UpdatedAtLte)
}

func mapUpdateGroupParams(params UpdateParams) (ugp sqlc.UpdateGroupParams) {
	ugp.AccountID = uuid.MustParse(params.AccountID.TypeID.UUID())
	ugp.Description = sql.NullString(params.Description)
	ugp.Name = sql.NullString(params.Name)
	ugp.ParentGroupID = id.ToNullUUID(params.ParentGroup)
	return
}

// func mapListGroupsParams(params ListGroupsParams, accountID typeid.TypeID) (sqlc.ListGroupsParams, error) {
// 	var lgp sqlc.ListGroupsParams
// 	parentGroupID, err := com.FromStringPtr(params.ParentGroup)
// 	if err != nil {
// 		return lgp, err
// 	}
// 	lgp.AccountID = uuid.MustParse(accountID.UUID())
// 	lgp.Name = sql.NullString(str.ParsePtr(params.Name))
// 	lgp.ParentGroupID = com.ToNullUUID(parentGroupID)
// 	lgp.Description = sql.NullString(str.ParsePtr(params.Description))

// 	// MapTimeRange(list.RequestParams.CreatedAt, &lgp.CreatedAtGt, &lgp.CreatedAtGte, &lgp.CreatedAtLt, &lgp.CreatedAtLte)
// 	// MapTimeRange(list.RequestParams.UpdatedAt, &lgp.UpdatedAtGt, &lgp.UpdatedAtGte, &lgp.UpdatedAtLt, &lgp.UpdatedAtLte)
// 	// MapPaginationParams(*list.RequestParams.PaginationParams, &lgp)
// 	return lgp, nil
// }

// func mapUpdateGroupParams(params UpdateGroupParams, id, accountID typeid.TypeID) (sqlc.UpdateGroupParams, error) {
// 	var ugp sqlc.UpdateGroupParams
// 	parentGroupID, err := com.FromStringPtr(params.ParentGroup)
// 	if err != nil {
// 		return ugp, err
// 	}
// 	ugp.ID = uuid.MustParse(id.UUID())
// 	ugp.AccountID = uuid.MustParse(accountID.UUID())
// 	ugp.Name = sql.NullString(str.ParsePtr(params.Name))
// 	ugp.ParentGroupID = com.ToNullUUID(parentGroupID)
// 	ugp.Description = sql.NullString(str.ParsePtr(params.Description))
// 	return ugp, nil
// }
