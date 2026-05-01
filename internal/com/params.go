package com

import (
	"database/sql"
	"time"
)

type PaginationParams struct {
	EndingBefore  *string `json:"ending_before" validate:"omitnil,id"`
	StartingAfter *string `json:"starting_after" validate:"omitnil,id,excluded_with=EndingBefore"`
	Limit         *int    `json:"limit" validate:"omitnil,gt=0,lte=100"`
}

type TimeRange struct {
	Gt  *time.Time `json:"gt" validate:"omitnil"`
	Lt  *time.Time `json:"lt" validate:"omitnil"`
	Gte *time.Time `json:"gte" validate:"omitnil"`
	Lte *time.Time `json:"lte" validate:"omitnil"`
}

type Int32Range struct {
	Gt  *int32 `json:"gt" validate:"omitnil"`
	Lt  *int32 `json:"lt" validate:"omitnil"`
	Gte *int32 `json:"gte" validate:"omitnil"`
	Lte *int32 `json:"lte" validate:"omitnil"`
}

func MapTimeRange(src *TimeRange, gt, gte, lt, lte *sql.NullTime) {
	if src == nil {
		return
	}
	if src.Gt != nil {
		*gt = sql.NullTime{Time: *src.Gt, Valid: true}
	}
	if src.Gte != nil {
		*gte = sql.NullTime{Time: *src.Gte, Valid: true}
	}
	if src.Lt != nil {
		*lt = sql.NullTime{Time: *src.Lt, Valid: true}
	}
	if src.Lte != nil {
		*lte = sql.NullTime{Time: *src.Lte, Valid: true}
	}
}

func MapInt32Range(src *Int32Range, gt, gte, lt, lte *sql.NullInt32) {
	if src == nil {
		return
	}
	if src.Gt != nil {
		*gt = sql.NullInt32{Int32: *src.Gt, Valid: true}
	}
	if src.Gte != nil {
		*gte = sql.NullInt32{Int32: *src.Gte, Valid: true}
	}
	if src.Lt != nil {
		*lt = sql.NullInt32{Int32: *src.Lt, Valid: true}
	}
	if src.Lte != nil {
		*lte = sql.NullInt32{Int32: *src.Lte, Valid: true}
	}
}
