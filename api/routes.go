package api

import (
	"golang-log-server/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(handler *handlers.LogHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/log", handler.WriteLog).Methods("POST")
	router.NotFoundHandler = router.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()

	return router
}
