package server

import (
	"api/routers"
	"net/http"
)

func InitServer() {
	router := routers.Init()

	router.HandleFunc("/", okMessage).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func okMessage(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("All Ok"))
}