package models

import (
	"github.com/littletwolee/commons/mongo"
)

type base struct{}

func (b *base) IsNil(id mongo.ObjectID) bool {
	return id == nil
}
