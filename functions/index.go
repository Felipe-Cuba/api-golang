package functions

import (
	"api/functions/students"
	"api/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func GetStudent(response http.ResponseWriter, request *http.Request) {
	students := students.GetStudents()

	res, err := json.Marshal(students)
	
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	} 
	
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetStudentById(response http.ResponseWriter, request *http.Request) {
	idString := mux.Vars(request)["id"]

	id, err := strconv.Atoi(idString)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	student, err := students.GetStudentById(id)
	
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	} 

	res, err := json.Marshal(student)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func InsertStudent(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	
	var student models.Student
	err = json.Unmarshal(body, &student)
	
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	
	students.InsertStudent(student)
	
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Student inserted successfully!"))
}

func UpdateStudent(response http.ResponseWriter, request *http.Request) {
	idString := mux.Vars(request)["id"]

	id, err := strconv.Atoi(idString)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	body, err := io.ReadAll(request.Body)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	var student models.Student
	err = json.Unmarshal(body, &student)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	res, err := students.UpdateStudent(id, student)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
		
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(res))
}

func DeleteStudent(response http.ResponseWriter, request *http.Request) {
	idString := mux.Vars(request)["id"]

	id, err := strconv.Atoi(idString)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	
	res, err := students.DeleteStudent(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
		
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(res))
}

