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

type Entry struct {
	Mongo *commons.Mongo
}

func GetEntriesController(mongo *commons.Mongo) *Entry {
	return &Entry{
		Mongo: mongo,
	}
}

const (
	entries string = "entries"
)

func (e *Entry) GetEntry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
		return
	}
	q := make(map[string]interface{})
	q["_id"] = e.Mongo.ObjectIDHex(id)
	var entry models.Entry
	if _, err := e.Mongo.ViewOneC(entries, q, &entry); err != nil {
		utils.ServerError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, entry)); err != nil {
		utils.ServerError(w, err)
		return
	}
}

func (e *Entry) PostEntry(w http.ResponseWriter, r *http.Request) {
	var entry models.Entry
	if err := utils.HttpBodyUnmarshal(r.Body, &entry); err != nil {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
		return
	}
	entryPoint := &entry
	entryPoint.ID = e.Mongo.NewObjectID()
	entryPoint.Tags = e.setTags(entryPoint.ID.Hex(), entryPoint.Tags)
	id, err := e.Mongo.InsertC(entries, entryPoint)
	if err != nil {
		utils.ServerError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, id)); err != nil {
		utils.ServerError(w, err)
		return
	}
}

func (e *Entry) setTags(eID string, entries []string) []string {
	var ids []string
	if len(entries) > 0 {
		wg := &sync.WaitGroup{}
		for _, v := range entries {
			if v != "" {
				wg.Add(1)
				m := make(map[string]string)
				m["title"] = v
				var tag models.Tag
				_, err := e.Mongo.ViewOneC(tags, m, &tag)
				if err != nil {
					if err.Error() == utils.NOT_FOUND {
						tag = models.Tag{
							Title:   v,
							Entries: []string{eID},
						}
						tagID, err := e.Mongo.InsertC(tags, &tag)
						if err != nil {
							catchErr(wg, err)
							continue
						}
						ids = append(ids, tagID)
					} else {
						catchErr(wg, err)
						continue
					}
				} else {
					m := make(map[string]interface{})
					m["entries"] = append(tag.Entries, eID)
					q := make(map[string]interface{})
					q["_id"] = tag.ID
					_, err := e.Mongo.UpdateC(tags, q, m)
					if err != nil {
						catchErr(wg, err)
						continue
					}
					ids = append(ids, tag.ID.Hex())
				}
				wg.Done()
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
