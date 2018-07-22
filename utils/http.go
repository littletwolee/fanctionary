package utils

import (
	"encoding/json"
	"fullday/models"
	"net/http"
)

const (
	ERROR_HTTP_BAD_REQUEST = "bad request"
)

func BadRequest(w http.ResponseWriter, err error) {
	serveError(w, http.StatusBadRequest, err)
}

func ServerError(w http.ResponseWriter, err error) {
	serveError(w, http.StatusInternalServerError, err)
}

func serveError(w http.ResponseWriter, code int, err error) {
	res := models.Result{}
	res.Msg = err.Error()
	jsonResult, _ := json.Marshal(res)
	w.WriteHeader(code)
	w.Write(jsonResult)
}
