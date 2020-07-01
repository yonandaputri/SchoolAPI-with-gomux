package teacherRepository

import "mysqlApp/api/master/models"

type TeacherRepository interface {
	GetAllTeacher() ([]*models.Teacher, error)
	GetTeacherById(id int) (*models.Teacher, error)
	AddTeacher(teacher *models.Teacher) error
	UpdateTeacher(id int, teacher *models.Teacher) error
	DeleteTeacher(id int) error
}
