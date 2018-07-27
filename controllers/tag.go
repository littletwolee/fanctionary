package controllers

import (
	"encoding/json"
	"fanctionary/models"
	"fanctionary/utils"
	"fmt"
	"net/http"

	"github.com/littletwolee/commons"

	"github.com/gorilla/mux"
)

const (
	tagsTable string = "tags"
)

func (s *Server) GetTag(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
		return
	}
	var tag models.Tag
	if _, err := s.Mongo.ViewOneC(tagsTable, s.Mongo.NewQuery("_id", s.Mongo.ObjectIDHex(id)), &tag); err != nil {
		utils.ServerError(w, err)
		return
	}
	tagAddr := &tag
	entries, err := s.getEntriesByIds(tagAddr.Entries)
	if err != nil {
		utils.ServerError(w, err)
		return
	}
	tagAddr.Entries = entries
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, tagAddr.IsNil())); err != nil {
		utils.ServerError(w, err)
		return
	}
}

func (s *Server) PostTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag
	if err := utils.HttpBodyUnmarshal(r.Body, &tag); err != nil {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
		return
	}
	id, err := s.Mongo.InsertC(tagsTable, tag)
	if err != nil {
		utils.ServerError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, id)); err != nil {
		utils.ServerError(w, err)
		return
	}
}

func (s *Server) getTagsByIds(tagIds []models.Tag) ([]models.Tag, error) {
	var ids commons.ObjectIDs
	for _, tag := range tagIds {
		ids = append(ids, tag.ID)
	}
	var tags []models.Tag
	if _, err := s.Mongo.ViewAllC(tagsTable,
		s.Mongo.In("_id", ids),
		&tags,
		s.Mongo.Select([]string{"_id", "title"})); err != nil {
		return nil, err
	}
	return tags, nil
}

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
