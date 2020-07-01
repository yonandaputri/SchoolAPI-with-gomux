package studentRepository

import "mysqlApp/api/master/models"

type StudentRepository interface {
	GetAllStudent() ([]*models.Student, error)
	GetStudentById(id int) (*models.Student, error)
	AddStudent(student *models.Student) error
	UpdateStudent(id int, student *models.Student) error
	DeleteStudent(id int) error
}
