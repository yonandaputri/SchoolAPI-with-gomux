package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mysqlApp/api/master/models"
	"mysqlApp/api/master/models/myResponse"
	"mysqlApp/api/master/usecases/teacherUsecase"
	"mysqlApp/utils"
	"net/http"
	"strconv"
)

type TeacherHandler struct {
	TeacherUseCase teacherUsecase.TeacherUseCase
}

func TeacherController(r *mux.Router, service teacherUsecase.TeacherUseCase) {
	teacherHandler := TeacherHandler{service}
	r.HandleFunc("/teachers", teacherHandler.ListTeachers).Methods(http.MethodGet)
	r.HandleFunc("/teacher/{id}", teacherHandler.GetTeacherById).Methods(http.MethodGet)
	r.HandleFunc("/teacher", teacherHandler.PostTeacher).Methods(http.MethodPost)
	r.HandleFunc("/teacher/{id}", teacherHandler.PutTeacher).Methods(http.MethodPut)
	r.HandleFunc("/teacher/{id}", teacherHandler.DeleteTeacher).Methods(http.MethodDelete)
}

func (t TeacherHandler) ListTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := t.TeacherUseCase.GetTeachers()

	var response myResponse.MessageResponse
	response.Message = "Data Teacher berhasil diambil"
	response.Data = teachers

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfTeachers, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfTeachers))
}

func (t TeacherHandler) GetTeacherById(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idTeacher, _ := strconv.Atoi(id)
	teacher, err := t.TeacherUseCase.GetTeacher(idTeacher)

	var response myResponse.MessageResponse
	response.Message = "Data Teacher berhasil diambil"
	response.Data = teacher

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfTeacher, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfTeacher))
}

func (t TeacherHandler) PostTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher []*models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	for _, value := range teacher {
		err := t.TeacherUseCase.PostTeacher(value)
		if err != nil {
			w.Write([]byte("Cannot add data"))
		}
	}
	w.Write([]byte("Succes add data"))
}

func (t TeacherHandler) PutTeacher(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idTeacher, _ := strconv.Atoi(id)
	var teacher *models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	err = t.TeacherUseCase.PutTeacher(idTeacher, teacher)
	if err != nil {
		w.Write([]byte("Cannot update data"))
	}
	w.Write([]byte("Succes update data"))
}

func (t TeacherHandler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idTeacher, _ := strconv.Atoi(id)
	err := t.TeacherUseCase.DeleteTeacher(idTeacher)
	if err != nil {
		w.Write([]byte("Cannot delete data"))
	}
	w.Write([]byte("Succes delete data"))
}
