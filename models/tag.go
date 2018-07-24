package models

import (
	"time"
)

type Tag struct {
	ID      string    `json:"_id" bson:"_id"`
	CTime   time.Time `json:"ctime" bson:"ctime"`
	Title   string    `json:"title" bson:"title"`
	Entries []string  `json:"entries" bson:"entries"`
}
