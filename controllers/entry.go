package controllers

import (
	"fmt"
	"net/http"
	"time"

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
	if err := e.Mongo.ViewOneC(entries, q, entry); err != nil {
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
	// entry.ID = bson.NewObjectId().String()
	entry.CTime = time.Now()
	if err := e.Mongo.InsertC(entries, &entry); err != nil {
		utils.ServerError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(models.NewResult(nil, nil)); err != nil {
		utils.ServerError(w, err)
		return
	}
}
