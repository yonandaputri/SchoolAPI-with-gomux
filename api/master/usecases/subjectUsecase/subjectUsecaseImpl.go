package subjectUsecase

import (
	"mysqlApp/api/master/models"
	"mysqlApp/api/master/repositories/subjectRepository"
)

type SubjectUsecaseImpl struct {
	subjectRepo subjectRepository.SubjectRepository
}

func (s SubjectUsecaseImpl) GetSubjects() ([]*models.Subject, error) {
	subjects, err := s.subjectRepo.GetAllSubject()
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

func (s SubjectUsecaseImpl) GetSubject(id int) (*models.Subject, error) {
	subject, err := s.subjectRepo.GetSubjectById(id)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (s SubjectUsecaseImpl) PostSubject(subject *models.Subject) error {
	err := s.subjectRepo.AddSubject(subject)
	if err != nil {
		return err
	}
	return nil
}

func (s SubjectUsecaseImpl) PutSubject(id int, subject *models.Subject) error {
	err := s.subjectRepo.UpdateSubject(id, subject)
	if err != nil {
		return err
	}
	return nil
}

func (s SubjectUsecaseImpl) DeleteSubject(id int) error {
	err := s.subjectRepo.DeleteSubject(id)
	if err != nil {
		return err
	}
	return nil
}

func InitSubjectUseCase(subjectRepo subjectRepository.SubjectRepository) SubjectUseCase {
	return &SubjectUsecaseImpl{subjectRepo}
}
