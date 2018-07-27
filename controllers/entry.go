package controllers

import (
	"fmt"
	"net/http"
	"sync"

	"fanctionary/models"
	"fanctionary/utils"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/littletwolee/commons"
)

const (
	entriesTable string = "entries"
)

func (s *Server) GetEntry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
		return
	}
	var entry models.Entry
	if _, err := s.Mongo.ViewOneC(entriesTable,
		s.Mongo.NewQuery("_id", s.Mongo.ObjectIDHex(id)),
		&entry); err != nil {
		utils.ServerError(w, err)
		return
	}
	entryAddr := &entry
	tags, err := s.getTagsByIds(entryAddr.Tags)
	if err != nil {
		utils.ServerError(w, err)
		return
	}
	entryAddr.Tags = tags
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, entryAddr.IsNil())); err != nil {
		utils.ServerError(w, err)
		return
	}
}

func (s *Server) PostEntry(w http.ResponseWriter, r *http.Request) {
	var entry models.Entry
	if err := utils.HttpBodyUnmarshal(r.Body, &entry); err != nil {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
		return
	}
	entryPoint := &entry
	entryPoint.ID = s.Mongo.NewObjectID()
	entryPoint.Tags = s.setTags(entryPoint.ID.Hex(), entryPoint.Tags)
	id, err := s.Mongo.InsertC(entriesTable, entryPoint)
	if err != nil {
		utils.ServerError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, id)); err != nil {
		utils.ServerError(w, err)
		return
	}
}
func (s *Server) getEntriesByIds(entryIds []models.Entry) ([]models.Entry, error) {
	var ids commons.ObjectIDs
	for _, entry := range entryIds {
		ids = append(ids, entry.ID)
	}
	var entries []models.Entry
	if _, err := s.Mongo.ViewAllC(entriesTable,
		s.Mongo.In("_id", ids),
		&entries,
		s.Mongo.Select([]string{"_id", "title"})); err != nil {
		return nil, err
	}
	return entries, nil
}
func (s *Server) setTags(eID string, tags []models.Tag) []models.Tag {
	eObjID := s.Mongo.ObjectIDHex(eID)
	var ids []models.Tag
	if len(tags) > 0 {
		wg := &sync.WaitGroup{}
		for _, v := range tags {
			if v.Title != "" {
				wg.Add(1)
				go func(wg *sync.WaitGroup, ids *[]models.Tag, v models.Tag) {
					m := s.Mongo.NewQuery("title", v.Title)
					var tag models.Tag
					_, err := s.Mongo.ViewOneC(tagsTable, m, &tag)
					if err != nil {
						if err == commons.ErrNotFound {
							tag = models.Tag{
								Title: v.Title,
								Entries: []models.Entry{
									models.Entry{ID: eObjID},
								},
							}
							tagID, err := s.Mongo.InsertC(tagsTable, &tag)
							if err != nil {
								catchErr(wg, err)
								return
							}
							if id := s.Mongo.ObjectIDHex(tagID); id != nil {
								*ids = append(*ids, models.Tag{ID: id})
							}
						} else {
							catchErr(wg, err)
							return
						}
					} else {
						m := s.Mongo.NewQuery("entries", append(tag.Entries, models.Entry{ID: eObjID}))
						q := s.Mongo.NewQuery("_id", tag.ID)
						_, err := s.Mongo.UpdateC(tagsTable, q, m)
						if err != nil {
							catchErr(wg, err)
							return
						}
						*ids = append(*ids, models.Tag{ID: tag.ID})
					}
					wg.Done()
				}(wg, &ids, v)
			}
		}
		wg.Wait()
	}
	return ids
}

func catchErr(wg *sync.WaitGroup, err error) {
	commons.Console().Error(err)
	wg.Done()
}
