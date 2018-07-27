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
	server := &controllers.Server{Mongo: r.Mongo}
	routerEntries := "/entries"
	router.HandleFunc(routerEntries, server.PostEntry).Methods(http.MethodPost)
	routerEntry := "/entry"
	router.HandleFunc(fmt.Sprintf("%s/%s", routerEntry, ID), server.GetEntry).Methods(http.MethodGet)

	routerTags := "/tags"
	router.HandleFunc(routerTags, server.PostTag).Methods(http.MethodPost)
	routerTag := "/tag"
	router.HandleFunc(fmt.Sprintf("%s/%s", routerTag, ID), server.GetTag).Methods(http.MethodGet)
}
