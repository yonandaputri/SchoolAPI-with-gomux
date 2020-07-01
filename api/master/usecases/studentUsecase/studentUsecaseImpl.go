package studentUsecase

import (
	"mysqlApp/api/master/models"
	"mysqlApp/api/master/repositories/studentRepository"
)

type StudentUsecaseImpl struct {
	studentRepo studentRepository.StudentRepository
}

func (s StudentUsecaseImpl) GetStudents() ([]*models.Student, error) {
	students, err := s.studentRepo.GetAllStudent()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s StudentUsecaseImpl) GetStudent(id int) (*models.Student, error) {
	student, err := s.studentRepo.GetStudentById(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s StudentUsecaseImpl) PostStudent(student *models.Student) error {
	err := s.studentRepo.AddStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (s StudentUsecaseImpl) PutStudent(id int, student *models.Student) error {
	err := s.studentRepo.UpdateStudent(id, student)
	if err != nil {
		return err
	}
	return nil
}

func (s StudentUsecaseImpl) DeleteStudent(id int) error {
	err := s.studentRepo.DeleteStudent(id)
	if err != nil {
		return err
	}
	return nil
}

func InitStudentUseCase(studentRepo studentRepository.StudentRepository) StudentUseCase {
	return &StudentUsecaseImpl{studentRepo}
}
