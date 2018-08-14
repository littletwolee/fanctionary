package models

import (
	"time"

	"github.com/littletwolee/commons/mongo"
)

type Entry struct {
	ID          mongo.ObjectID `json:"_id" bson:"_id"`
	CTime       time.Time      `json:"-" bson:"ctime,omitempty"`
	Title       string         `json:"title,omitempty" bson:"title,omitempty"`
	Explanation string         `json:"exp,omitempty" bson:"exp,omitempty"`
	Tags        []Tag          `json:"tags,omitempty" bson:"tags,omitempty"`
}

// func (e *Entry) SetTags(f func(tags []string) []string) []string {
// 	if len(e.Tags) > 0 {
// 		return f(e.Tags)
// 	}
// 	return nil
// }

func (e *Entry) IsNil() interface{} {
	if e == nil || e.ID == nil {
		return nil
	}
	return e
}
