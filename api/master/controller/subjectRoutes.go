package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mysqlApp/api/master/models"
	"mysqlApp/api/master/models/myResponse"
	"mysqlApp/api/master/usecases/subjectUsecase"
	"mysqlApp/utils"
	"net/http"
	"strconv"
)

type SubjectHandler struct {
	SubjectUseCase subjectUsecase.SubjectUseCase
}

func SubjectController(r *mux.Router, service subjectUsecase.SubjectUseCase) {
	subjectHandler := SubjectHandler{service}
	r.HandleFunc("/subjects", subjectHandler.ListSubjects).Methods(http.MethodGet)
	r.HandleFunc("/subject/{id}", subjectHandler.GetSubjectById).Methods(http.MethodGet)
	r.HandleFunc("/subject", subjectHandler.PostSubject).Methods(http.MethodPost)
	r.HandleFunc("/subject/{id}", subjectHandler.PutSubject).Methods(http.MethodPut)
	r.HandleFunc("/subject/{id}", subjectHandler.DeleteSubject).Methods(http.MethodDelete)
}

func (s SubjectHandler) ListSubjects(w http.ResponseWriter, r *http.Request) {
	subjects, err := s.SubjectUseCase.GetSubjects()

	var response myResponse.MessageResponse
	response.Message = "Data Subject berhasil diambil"
	response.Data = subjects

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfSubjects, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfSubjects))
}

func (s SubjectHandler) GetSubjectById(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idSubject, _ := strconv.Atoi(id)
	subject, err := s.SubjectUseCase.GetSubject(idSubject)

	var response myResponse.MessageResponse
	response.Message = "Data Subject berhasil diambil"
	response.Data = subject

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfSubject, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfSubject))
}

func (s SubjectHandler) PostSubject(w http.ResponseWriter, r *http.Request) {
	var subject []*models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	for _, value := range subject {
		err := s.SubjectUseCase.PostSubject(value)
		if err != nil {
			w.Write([]byte("Cannot add data"))
		}
	}
	w.Write([]byte("Succes add data"))
}

func (s SubjectHandler) PutSubject(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idSubject, _ := strconv.Atoi(id)
	var subject *models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	err = s.SubjectUseCase.PutSubject(idSubject, subject)
	if err != nil {
		w.Write([]byte("Cannot update data"))
	}
	w.Write([]byte("Succes update data"))
}

func (s SubjectHandler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := utils.GetPathVar(key, r)
	idSubject, _ := strconv.Atoi(id)
	err := s.SubjectUseCase.DeleteSubject(idSubject)
	if err != nil {
		w.Write([]byte("Cannot delete data"))
	}
	w.Write([]byte("Succes delete data"))
}
