package groups

import "github.com/d-darac/lagra/internal/com"

type CreateGroupParams struct {
	Description *string  `json:"description" validate:"omitnil"`
	Name        string   `json:"name" validate:"required"`
	ParentGroup *string  `json:"parent_group" validate:"omitnil,id"`
	Expand      []string `json:"expand" validate:"omitnil,dive,oneof=parent_group"`
}

type GetGroupParams struct {
	Expand []string `json:"expand" validate:"omitnil,dive,oneof=parent_group"`
}

type ListGroupsParams struct {
	*com.PaginationParams
	CreatedAt   *com.TimeRange `json:"created_at" validate:"omitnil"`
	UpdatedAt   *com.TimeRange `json:"updated_at" validate:"omitnil"`
	Description *string        `json:"description" validate:"omitnil"`
	Name        *string        `json:"name" validate:"omitnil"`
	ParentGroup *string        `json:"parent_group" validate:"omitnil,id"`
	Expand      []string       `json:"expand" validate:"omitnil,dive,oneof=parent_group"`
}

type UpdateGroupParams struct {
	Description *string  `json:"description" validate:"omitnil"`
	Name        *string  `json:"name" validate:"omitnil"`
	ParentGroup *string  `json:"parent_group" validate:"omitnil,id"`
	Expand      []string `json:"expand" validate:"omitnil,dive,oneof=parent_group"`
}
