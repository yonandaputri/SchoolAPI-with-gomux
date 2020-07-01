package master

import (
	"database/sql"
	"mysqlApp/api/master/controller"
	"mysqlApp/api/master/repositories/studentRepository"
	"mysqlApp/api/master/repositories/subjectRepository"
	"mysqlApp/api/master/repositories/teacherRepository"
	"mysqlApp/api/master/usecases/studentUsecase"
	"mysqlApp/api/master/usecases/subjectUsecase"
	"mysqlApp/api/master/usecases/teacherUsecase"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	studentRepo := studentRepository.InitStudentRepoImpl(db)
	studentUseCase := studentUsecase.InitStudentUseCase(studentRepo)
	teacherRepo := teacherRepository.InitTeacherRepoImpl(db)
	teacherUseCase := teacherUsecase.InitTeacherUseCase(teacherRepo)
	subjectRepo := subjectRepository.InitSubjectRepoImpl(db)
	subjectUseCase := subjectUsecase.InitSubjectUseCase(subjectRepo)
	controller.StudentController(r, studentUseCase)
	controller.TeacherController(r, teacherUseCase)
	controller.SubjectController(r, subjectUseCase)
}
