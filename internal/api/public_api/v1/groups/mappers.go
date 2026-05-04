package groups

import (
	"time"

	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/services/groups"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/d-darac/lagra/pkg/str"
	"github.com/d-darac/lagra/pkg/t"
)

func mapCreateParams(params CreateGroupParams, accountID id.ID) (cp groups.CreateParams, err error) {
	cp.ParentGroup, err = id.ParseNull(params.ParentGroup)
	if err != nil {
		return
	}
	cp.AccountID = accountID
	cp.GroupID = id.Generate(string(com.IDPrefixAccount))
	cp.CreatedAt = cp.GroupID.Time()
	cp.UpdatedAt = cp.GroupID.Time()
	cp.Description = str.ParseNull(params.Description)
	cp.Name = params.Name
	return
}

func mapUpdateParams(params UpdateGroupParams, groupID, accountID id.ID) (up groups.UpdateParams, err error) {
	up.ParentGroup, err = id.ParseNull(params.ParentGroup)
	if err != nil {
		return
	}
	up.AccountID = accountID
	up.GroupID = groupID
	up.Description = str.ParseNull(params.Description)
	up.Name = str.ParseNull(params.Name)
	up.UpdatedAt = time.Now().UTC()
	return
}

func mapListParams(params ListGroupsParams, accountID id.ID) (lp groups.ListParams, err error) {
	lp.EndingBefore, err = id.ParseNull(params.EndingBefore)
	if err != nil {
		return
	}
	lp.StartingAfter, err = id.ParseNull(params.StartingAfter)
	if err != nil {
		return
	}
	lp.AccountID = accountID
	lp.CreatedAtGt = t.ParseNull(params.CreatedAt.Gt)
	lp.CreatedAtGte = t.ParseNull(params.CreatedAt.Gte)
	lp.CreatedAtLt = t.ParseNull(params.CreatedAt.Lt)
	lp.CreatedAtLte = t.ParseNull(params.CreatedAt.Lte)
	lp.Description = str.ParseNull(params.Description)
	lp.Name = str.ParseNull(params.Name)
	lp.Limit = i32.ParseNull(params.Limit)
	lp.UpdatedAtGt = t.ParseNull(params.UpdatedAt.Gt)
	lp.UpdatedAtGte = t.ParseNull(params.UpdatedAt.Gte)
	lp.UpdatedAtLt = t.ParseNull(params.UpdatedAt.Lt)
	lp.UpdatedAtLte = t.ParseNull(params.UpdatedAt.Lte)
	return
}
