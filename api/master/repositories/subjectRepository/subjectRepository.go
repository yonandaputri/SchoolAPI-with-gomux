package subjectRepository

import "mysqlApp/api/master/models"

type SubjectRepository interface {
	GetAllSubject() ([]*models.Subject, error)
	GetSubjectById(id int) (*models.Subject, error)
	AddSubject(subject *models.Subject) error
	UpdateSubject(id int, subject *models.Subject) error
	DeleteSubject(id int) error
}
