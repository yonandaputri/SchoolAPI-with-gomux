package controller

import (
	"encoding/json"
	"mysqlApp/api/master/models"
	"mysqlApp/api/master/models/myResponse"
	"mysqlApp/api/master/usecases/studentUsecase"
	"mysqlApp/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	StudentUseCase studentUsecase.StudentUseCase
}

func StudentController(r *mux.Router, service studentUsecase.StudentUseCase) {
	studentHandler := StudentHandler{service}
	r.HandleFunc("/students", studentHandler.ListStudents).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", studentHandler.GetStudentById).Methods(http.MethodGet)
	r.HandleFunc("/student", studentHandler.PostStudent).Methods(http.MethodPost)
	r.HandleFunc("/student/{id}", studentHandler.PutStudent).Methods(http.MethodPut)
	r.HandleFunc("/student/{id}", studentHandler.DeleteStudent).Methods(http.MethodDelete)
}

func (s StudentHandler) ListStudents(w http.ResponseWriter, r *http.Request) {
	students, err := s.StudentUseCase.GetStudents()

	var response myResponse.MessageResponse
	response.Message = "Data Student berhasil diambil"
	response.Data = students

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfStudents, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfStudents))
}

func (s StudentHandler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idStudent, _ := strconv.Atoi(id)
	student, err := s.StudentUseCase.GetStudent(idStudent)

	var response myResponse.MessageResponse
	response.Message = "Data Student berhasil diambil"
	response.Data = student

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfStudents, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfStudents))
}

func (s StudentHandler) PostStudent(w http.ResponseWriter, r *http.Request) {
	var student []*models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	for _, value := range student {
		err := s.StudentUseCase.PostStudent(value)
		if err != nil {
			w.Write([]byte("Cannot add data"))
		}
	}
	w.Write([]byte("Succes add data"))
}

func (s StudentHandler) PutStudent(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idStudent, _ := strconv.Atoi(id)
	var student *models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	err = s.StudentUseCase.PutStudent(idStudent, student)
	if err != nil {
		w.Write([]byte("Cannot update data"))
	}
	w.Write([]byte("Succes update data"))
}

func (s StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idStudent, _ := strconv.Atoi(id)
	err := s.StudentUseCase.DeleteStudent(idStudent)
	if err != nil {
		w.Write([]byte("Cannot delete data"))
	}
	w.Write([]byte("Succes delete data"))
}
