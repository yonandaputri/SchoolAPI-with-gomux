package teacherRepository

import (
	"database/sql"
	"mysqlApp/api/master/models"
)

type TeacherRepoImpl struct {
	db *sql.DB
}

func (t TeacherRepoImpl) GetAllTeacher() ([]*models.Teacher, error) {
	query := "SELECT * FROM teacher"
	data, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	var teachers []*models.Teacher

	for data.Next() {
		var teacher = new(models.Teacher)
		var err = data.Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.Email)

		if err != nil {
			return nil, err
		}

		teachers = append(teachers, teacher)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (t TeacherRepoImpl) GetTeacherById(id int) (*models.Teacher, error) {
	var teacher models.Teacher
	query := "SELECT * FROM teacher WHERE id=?"
	err := t.db.QueryRow(query, id).Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.Email)

	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (t TeacherRepoImpl) AddTeacher(teacher *models.Teacher) error {
	data, err := t.db.Begin()

	if err != nil {
		return err
	}

	query := "INSERT INTO teacher(first_name,last_name,email) VALUES (?, ?, ?)"
	row, err := t.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = row.Exec(teacher.FirstName, teacher.FirstName, teacher.Email)
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

func (t TeacherRepoImpl) UpdateTeacher(id int, teacher *models.Teacher) error {
	data, err := t.db.Begin()
	if err != nil {
		return err
	}

	_, _ = t.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	query := "UPDATE teacher SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	row, err := t.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = row.Exec(teacher.FirstName, teacher.LastName, teacher.Email, id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	_, _ = t.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}

func (t TeacherRepoImpl) DeleteTeacher(id int) error {
	data, err := t.db.Begin()
	if err != nil {
		return err
	}

	_, _ = t.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	query := "DELETE FROM teacher WHERE id = ?"
	row, err := t.db.Prepare(query)
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
	_, _ = t.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}

func InitTeacherRepoImpl(db *sql.DB) TeacherRepository {
	return &TeacherRepoImpl{db}
}
