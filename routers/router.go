package routers

import (
	"fanctionary/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct{}

func (r *Router) Handler() http.Handler {
	router := mux.NewRouter()
	r.handleAPI(router)
	return router
}

var advice = controllers.GetAdviceController()

func (r *Router) handleAPI(router *mux.Router) {
	routerAdvice := "/advice"
	router.HandleFunc(fmt.Sprintf("%s/{open_id}/type/{type}", routerAdvice), advice.GetAdvice).Methods(http.MethodGet)
}
