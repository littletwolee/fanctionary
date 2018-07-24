package controllers

import (
	"github.com/littletwolee/commons"
)

type Tag struct {
	Mongo *commons.Mongo
}

func GetTagsController(mongo *commons.Mongo) *Tag {
	return &Tag{
		Mongo: mongo,
	}
}

const (
	tags string = "tags"
)

// func (t *Tag) GetTag(w http.ResponseWriter, r *http.Request) {
// 	ID := mux.Vars(r)["id"]
// 	if ID == "" {
// 		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
// 		return
// 	}
// 	q := make(map[string]interface{})
// 	q["_id"] = ID
// 	var tag models.Tag
// 	if err := t.Mongo.ViewOneC(tags, q, tag); err != nil {
// 		utils.ServerError(w, err)
// 		return
// 	}
// 	if err := json.NewEncoder(w).Encode(models.NewResult(nil, tag)); err != nil {
// 		utils.ServerError(w, err)
// 		return
// 	}
// }

// func (t *Tag) getTag(tagStr string) (*models.Tag, error) {
// 	q := make(map[string]interface{})
// 	q["tag"] = tagStr
// 	var tag *models.Tag
// 	if err := t.Mongo.ViewOneC(tags, q, *tag); err != nil {
// 		return nil, err
// 	}
// 	return tag, nil
// }

// func (t *Tag) PostTag(w http.ResponseWriter, r *http.Request) {
// 	var tag models.Tag
// 	if err := utils.HttpBodyUnmarshal(r.Body, &tag); err != nil {
// 		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
// 		return
// 	}
// 	if err := t.Mongo.InsertC(tags, &tag); err != nil {
// 		utils.ServerError(w, err)
// 		return
// 	}
// 	if err := json.NewEncoder(w).Encode(models.NewResult(nil, nil)); err != nil {
// 		utils.ServerError(w, err)
// 		return
// 	}
// }
