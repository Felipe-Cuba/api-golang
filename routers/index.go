package routers

import "github.com/gorilla/mux"

var router *mux.Router

func Init() *mux.Router {
	router = mux.NewRouter()

	applyRoutes(router)

	return router
}

func GetRouter() *mux.Router {
	return router
}