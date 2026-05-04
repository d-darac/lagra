package groups

import (
	"time"

	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/d-darac/lagra/pkg/str"
	"github.com/d-darac/lagra/pkg/t"
)

type CreateParams struct {
	AccountID, GroupID   id.ID
	CreatedAt, UpdatedAt time.Time
	Description          str.NullString
	Name                 string
	ParentGroup          id.NullID
}

type DeleteParams struct {
	AccountID, GroupID id.ID
}

type GetParams struct {
	AccountID, GroupID id.ID
}

type ListParams struct {
	CreatedAtGt, CreatedAtGte,
	CreatedAtLt, CreatedAtLte,
	UpdatedAtGt, UpdatedAtGte,
	UpdatedAtLt, UpdatedAtLte t.NullTime
	AccountID                                id.ID
	Description, Name                        str.NullString
	EndingBefore, StartingAfter, ParentGroup id.NullID
	Limit                                    i32.NullInt32
}

type GetByIDsParams struct {
	AccountID id.ID
	GroupIDs  []id.ID
}

type UpdateParams struct {
	AccountID, GroupID id.ID
	UpdatedAt          time.Time
	Description, Name  str.NullString
	ParentGroup        id.NullID
}
