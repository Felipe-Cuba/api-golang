package routers

import (
	"api/functions"

	"github.com/gorilla/mux"
)

func applyRoutes(router *mux.Router) {
	router.HandleFunc("/student", functions.GetStudent).Methods("GET")
	router.HandleFunc("/student", functions.InsertStudent).Methods("POST")
	router.HandleFunc("/student/{id}", functions.GetStudentById).Methods("GET")
	router.HandleFunc("/student/{id}", functions.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{id}", functions.DeleteStudent).Methods("DELETE")
}