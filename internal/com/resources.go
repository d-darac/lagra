package com

import "go.jetify.com/typeid/v2"

type Resource interface {
	AccountID() typeid.TypeID
}
