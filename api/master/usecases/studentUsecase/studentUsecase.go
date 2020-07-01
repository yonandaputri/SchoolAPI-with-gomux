package studentUsecase

import "mysqlApp/api/master/models"

type StudentUseCase interface {
	GetStudents() ([]*models.Student, error)
	GetStudent(id int) (*models.Student, error)
	PostStudent(student *models.Student) error
	PutStudent(id int, student *models.Student) error
	DeleteStudent(id int) error
}
