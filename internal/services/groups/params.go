package groups

import (
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/d-darac/lagra/pkg/str"
	"github.com/d-darac/lagra/pkg/t"
)

type CreateParams struct {
	AccountID   id.ID
	Description str.NullString
	Name        string
	ParentGroup id.NullID
}

type DeleteParams struct {
	AccountID, GroupID id.ID
}

type GetParams struct {
	AccountID, GroupID id.ID
}

type ListParams struct {
	CreatedAtGt   t.NullTime
	CreatedAtGte  t.NullTime
	CreatedAtLt   t.NullTime
	CreatedAtLte  t.NullTime
	UpdatedAtGt   t.NullTime
	UpdatedAtGte  t.NullTime
	UpdatedAtLt   t.NullTime
	UpdatedAtLte  t.NullTime
	AccountID     id.ID
	Description   str.NullString
	Name          str.NullString
	EndingBefore  id.NullID
	StartingAfter id.NullID
	ParentGroup   id.NullID
	Limit         i32.NullInt32
}

type GetByIDsParams struct {
	AccountID id.ID
	IDs       []id.ID
}

type UpdateParams struct {
	AccountID, GroupID id.ID
	Description        str.NullString
	Name               str.NullString
	ParentGroup        id.NullID
}
