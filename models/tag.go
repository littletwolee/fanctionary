package models

import (
	"time"

	"github.com/littletwolee/commons"
)

type Tag struct {
	ID      commons.ObjectID `json:"_id" bson:"_id"`
	CTime   time.Time        `json:"-" bson:"ctime,omitempty"`
	Title   string           `json:"title,omitempty" bson:"title,omitempty"`
	Entries []Entry          `json:"entries,omitempty" bson:"entries,omitempty"`
}

func (t *Tag) IsNil() interface{} {
	if t == nil || t.ID == nil {
		return nil
	}
	return t
}
