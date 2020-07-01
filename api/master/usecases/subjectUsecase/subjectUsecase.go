package subjectUsecase

import "mysqlApp/api/master/models"

type SubjectUseCase interface {
	GetSubjects() ([]*models.Subject, error)
	GetSubject(id int) (*models.Subject, error)
	PostSubject(subject *models.Subject) error
	PutSubject(id int, subject *models.Subject) error
	DeleteSubject(id int) error
}
