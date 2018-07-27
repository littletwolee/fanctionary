package models

import (
	"github.com/littletwolee/commons"
)

type base struct{}

func (b *base) IsNil(id commons.ObjectID) bool {
	return id == nil
}
