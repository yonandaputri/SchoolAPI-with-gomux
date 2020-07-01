package studentRepository

import (
	"database/sql"
	"mysqlApp/api/master/models"
)

type StudentRepoImpl struct {
	db *sql.DB
}

func (s StudentRepoImpl) GetAllStudent() ([]*models.Student, error) {
	query := "SELECT * FROM student"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var students []*models.Student

	for data.Next() {
		var student = new(models.Student)
		var err = data.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return students, nil
	// panic("implement me")
}

func (s StudentRepoImpl) GetStudentById(id int) (*models.Student, error) {
	var student models.Student
	query := "SELECT * FROM student WHERE id=?"
	err := s.db.QueryRow(query, id).Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s StudentRepoImpl) AddStudent(student *models.Student) error {
	data, err := s.db.Begin()

	if err != nil {
		return err
	}

	query := "INSERT INTO student(first_name,last_name,email) VALUES (?, ?, ?)"
	row, err := s.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = row.Exec(student.FirstName, student.LastName, student.Email)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (s StudentRepoImpl) UpdateStudent(id int, student *models.Student) error {
	data, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	query := "UPDATE student SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	row, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = row.Exec(student.FirstName, student.LastName, student.Email, id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}

func (s StudentRepoImpl) DeleteStudent(id int) error {
	data, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	query := "DELETE FROM student WHERE id = ?"
	row, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = row.Exec(id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}

func InitStudentRepoImpl(db *sql.DB) StudentRepository {
	return &StudentRepoImpl{db}
}
