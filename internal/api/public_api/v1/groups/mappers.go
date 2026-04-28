package groups

import (
	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/services/groups"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/str"
	"go.jetify.com/typeid/v2"
)

func mapCreateParams(params CreateGroupParams, accountID typeid.TypeID) (cp groups.CreateParams, err error) {
	cp.ParentGroup, err = com.FromStringPtr(params.ParentGroup)
	if err != nil {
		return
	}
	cp.AccountID = accountID
	cp.Description = str.ParsePtr(params.Description)
	cp.Name = params.Name
	return
}

func mapUpdateParams(params UpdateGroupParams, groupID, accountID typeid.TypeID) (up groups.UpdateParams, err error) {
	up.ParentGroup, err = com.FromStringPtr(params.ParentGroup)
	if err != nil {
		return
	}
	up.AccountID = accountID
	up.GroupID = groupID
	up.Description = str.ParsePtr(params.Description)
	up.Name = str.ParsePtr(params.Name)
	return
}

func mapListParams(params ListGroupsParams, accountID typeid.TypeID) (lp groups.ListParams, err error) {
	lp.EndingBefore, err = com.FromStringPtr(params.EndingBefore)
	if err != nil {
		return
	}
	lp.StartingAfter, err = com.FromStringPtr(params.StartingAfter)
	if err != nil {
		return
	}
	lp.AccountID = accountID
	lp.Limit = i32.ParsePtr(params.Limit)
	return
}
