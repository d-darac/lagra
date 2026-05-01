package com

import "github.com/d-darac/lagra/pkg/id"

type Resource interface {
	AccountID() id.ID
}
