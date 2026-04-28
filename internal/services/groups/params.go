package groups

import (
	"database/sql"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/str"
	"go.jetify.com/typeid/v2"
)

type CreateParams struct {
	AccountID   typeid.TypeID
	Description str.NullString
	Name        string
	ParentGroup com.NullTypeID
}

type DeleteParams struct {
	AccountID, GroupID typeid.TypeID
}

type GetParams struct {
	AccountID, GroupID typeid.TypeID
}

type ListParams struct {
	AccountID                           typeid.TypeID
	EndingBefore, StartingAfter         com.NullTypeID
	startingAfterDate, endingBeforeDate sql.NullTime
	Limit                               i32.NullInt32
}

type GetByIDsParams struct {
	AccountID typeid.TypeID
	IDs       []typeid.TypeID
}

type UpdateParams struct {
	AccountID, GroupID typeid.TypeID
	Description        str.NullString
	Name               str.NullString
	ParentGroup        com.NullTypeID
}
