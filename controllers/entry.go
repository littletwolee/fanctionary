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
	ID := mux.Vars(r)["id"]
	if ID == "" {
		utils.BadRequest(w, fmt.Errorf(utils.ERROR_HTTP_BAD_REQUEST))
		return
	}
	q := make(map[string]interface{})
	q["_id"] = ID
	var entry models.Entry
	if _, err := e.Mongo.ViewOneC(entries, q, entry); err != nil {
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
	a := e.setTags(entry.Tags)
	fmt.Println(a)
	entry.Tags = a
	id, err := e.Mongo.InsertC(entries, &entry)
	if err != nil {
		utils.ServerError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, id)); err != nil {
		utils.ServerError(w, err)
		return
	}
}

func (e *Entry) setTags(entries []string) []string {
	var ids []string
	if len(entries) > 0 {
		var wg sync.WaitGroup
		for _, v := range entries {
			if v != "" {
				wg.Add(1)
				m := make(map[string][]string)
				m["tag"] = []string{v}
				id, err := e.Mongo.UpsertC(tags, m, &models.Tag{Title: v})
				if err != nil {
					commons.Console().Error(err)
					continue
				}
				ids = append(ids, id)
				wg.Done()
			}
		}
		wg.Wait()
	}
	return ids
}
