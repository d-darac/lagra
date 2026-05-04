package groups

import (
	"database/sql"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/google/uuid"
)

func mapCreateGroupParams(params CreateParams) (dbParams sqlc.CreateGroupParams) {
	dbParams.AccountID = params.AccountID.UUID()
	dbParams.CreatedAt = params.CreatedAt
	dbParams.Description = sql.NullString(params.Description)
	dbParams.ID = params.GroupID.UUID()
	dbParams.Name = params.Name
	dbParams.ParentGroupID = id.ToNullUUID(params.ParentGroup)
	dbParams.UpdatedAt = params.UpdatedAt
	return
}

func mapDeleteGroupParams(params DeleteParams) sqlc.DeleteGroupParams {
	return sqlc.DeleteGroupParams{
		AccountID: params.AccountID.UUID(),
		ID:        params.GroupID.UUID(),
	}
}

func mapGetGroupParams(params GetParams) sqlc.GetGroupParams {
	return sqlc.GetGroupParams{
		AccountID: params.AccountID.UUID(),
		ID:        params.GroupID.UUID(),
	}
}

func mapGetGroupsByIDsParams(params GetByIDsParams) (dbParams sqlc.GetGroupsByIDsParams) {
	dbParams.AccountID = params.AccountID.UUID()
	uuids := make([]uuid.UUID, len(params.GroupIDs))
	for i, ID := range params.GroupIDs {
		uuids[i] = ID.UUID()
	}
	dbParams.IDs = uuids
	return
}

func mapListGroupParams(params ListParams) (dbParams sqlc.ListGroupsParams) {
	dbParams.AccountID = params.AccountID.UUID()
	dbParams.Description = sql.NullString(params.Description)
	dbParams.Name = sql.NullString(params.Name)
	dbParams.ParentGroupID = id.ToNullUUID(params.ParentGroup)
	mapPaginationParams(params, &dbParams)
	mapTimeRangeParams(params, &dbParams)
	return
}

func mapPaginationParams(params ListParams, dbParams *sqlc.ListGroupsParams) {
	dbParams.EndingBefore = id.ToNullUUID(params.EndingBefore)
	dbParams.Limit = sql.NullInt32(params.Limit)
	dbParams.StartingAfter = id.ToNullUUID(params.StartingAfter)
}

func mapTimeRangeParams(params ListParams, dbParams *sqlc.ListGroupsParams) {
	dbParams.CreatedAtGt = sql.NullTime(params.CreatedAtGt)
	dbParams.CreatedAtGte = sql.NullTime(params.CreatedAtGte)
	dbParams.CreatedAtLt = sql.NullTime(params.CreatedAtLt)
	dbParams.CreatedAtLte = sql.NullTime(params.CreatedAtLte)
	dbParams.UpdatedAtGt = sql.NullTime(params.UpdatedAtGt)
	dbParams.UpdatedAtGte = sql.NullTime(params.UpdatedAtGte)
	dbParams.UpdatedAtLt = sql.NullTime(params.UpdatedAtLt)
	dbParams.UpdatedAtLte = sql.NullTime(params.UpdatedAtLte)
}

func mapUpdateGroupParams(params UpdateParams) (dbParams sqlc.UpdateGroupParams) {
	dbParams.AccountID = params.AccountID.UUID()
	dbParams.Description = sql.NullString(params.Description)
	dbParams.Name = sql.NullString(params.Name)
	dbParams.ParentGroupID = id.ToNullUUID(params.ParentGroup)
	return
}
