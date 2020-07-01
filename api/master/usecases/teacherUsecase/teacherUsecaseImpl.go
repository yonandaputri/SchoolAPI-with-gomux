package teacherUsecase

import (
	"mysqlApp/api/master/models"
	"mysqlApp/api/master/repositories/teacherRepository"
)

type TeacherUsecaseImpl struct {
	teacherRepo teacherRepository.TeacherRepository
}

func (t TeacherUsecaseImpl) GetTeachers() ([]*models.Teacher, error) {
	teachers, err := t.teacherRepo.GetAllTeacher()
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func (t TeacherUsecaseImpl) GetTeacher(id int) (*models.Teacher, error) {
	teacher, err := t.teacherRepo.GetTeacherById(id)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (t TeacherUsecaseImpl) PostTeacher(teacher *models.Teacher) error {
	err := t.teacherRepo.AddTeacher(teacher)
	if err != nil {
		return err
	}
	return nil
}

func (t TeacherUsecaseImpl) PutTeacher(id int, teacher *models.Teacher) error {
	err := t.teacherRepo.UpdateTeacher(id, teacher)
	if err != nil {
		return err
	}
	return nil
}

func (t TeacherUsecaseImpl) DeleteTeacher(id int) error {
	err := t.teacherRepo.DeleteTeacher(id)
	if err != nil {
		return err
	}
	return nil
}

func InitTeacherUseCase(teacherRepo teacherRepository.TeacherRepository) TeacherUseCase {
	return &TeacherUsecaseImpl{teacherRepo}
}
