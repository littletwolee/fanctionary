package routers

import (
	"fanctionary/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/littletwolee/commons"
)

const (
	ID = "{id}"
)

type Router struct {
	Mongo *commons.Mongo
}

func (r *Router) Handler() http.Handler {
	router := mux.NewRouter()
	r.handleAPI(router)
	return router
}

func (r *Router) handleAPI(router *mux.Router) {
	entry := controllers.GetEntriesController(r.Mongo)
	routerEntries := "/entries"
	router.HandleFunc(fmt.Sprintf("%s/%s", routerEntries, ID), entry.GetEntry).Methods(http.MethodGet)
	router.HandleFunc(routerEntries, entry.PostEntry).Methods(http.MethodPost)
}
