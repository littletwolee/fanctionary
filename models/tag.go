package models

import (
	"time"

	"github.com/littletwolee/commons"
)

type Tag struct {
	ID      commons.ObjectID `json:"_id" bson:"_id"`
	CTime   time.Time        `json:"ctime" bson:"ctime"`
	Title   string           `json:"title" bson:"title"`
	Entries []string         `json:"entries" bson:"entries"`
}
