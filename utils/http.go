package utils

import (
	"encoding/json"
	"fanctionary/models"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	ERROR_HTTP_BAD_REQUEST = "bad request"
	ERROR_SERVER_ERROR     = "server error"
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

func HttpBodyUnmarshal(r io.ReadCloser, object interface{}) error {
	defer r.Close()
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, object)
	if err != nil {
		return err
	}
	return nil
}
