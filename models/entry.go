package models

import (
	"time"
)

type Entry struct {
	ID          string    `json:"_id" bson:"_id"`
	CTime       time.Time `json:"ctime" bson:"ctime"`
	Title       string    `json:"title" bson:"title"`
	Explanation string    `json:"exp" bson:"exp"`
	Tags        []string  `json:"tags" bson:"tags"`
}

func (e *Entry) SetTags(f func(tags []string) []string) []string {
	if len(e.Tags) > 0 {
		return f(e.Tags)
	}
	return nil
}
