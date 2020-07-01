package teacherUsecase

import "mysqlApp/api/master/models"

type TeacherUseCase interface {
	GetTeachers() ([]*models.Teacher, error)
	GetTeacher(id int) (*models.Teacher, error)
	PostTeacher(teacher *models.Teacher) error
	PutTeacher(id int, teacher *models.Teacher) error
	DeleteTeacher(id int) error
}
