package groups

import (
	"github.com/d-darac/lagra/internal/services/groups"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/d-darac/lagra/pkg/str"
)

func mapCreateParams(params CreateGroupParams, accountID id.ID) (cp groups.CreateParams, err error) {
	cp.ParentGroup, err = id.ParseNull(params.ParentGroup)
	if err != nil {
		return
	}
	cp.AccountID = accountID
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
	lp.Limit = i32.ParseNull(params.Limit)
	return
}
