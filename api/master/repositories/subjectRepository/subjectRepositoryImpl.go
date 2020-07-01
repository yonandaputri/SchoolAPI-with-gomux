package subjectRepository

import (
	"database/sql"
	"mysqlApp/api/master/models"
)

type SubjectRepoImpl struct {
	db *sql.DB
}

func (s SubjectRepoImpl) GetAllSubject() ([]*models.Subject, error) {
	query := "SELECT * FROM subject"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	var subjects []*models.Subject

	for data.Next() {
		var subject = new(models.Subject)
		var err = data.Scan(&subject.Id, &subject.SubjectName)

		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s SubjectRepoImpl) GetSubjectById(id int) (*models.Subject, error) {
	var subject models.Subject
	query := "SELECT * FROM subject WHERE id=?"
	err := s.db.QueryRow(query, id).Scan(&subject.Id, &subject.SubjectName)

	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (s SubjectRepoImpl) AddSubject(subject *models.Subject) error {
	data, err := s.db.Begin()

	if err != nil {
		return err
	}

	query := "INSERT INTO subject(subject_name) VALUES (?)"
	row, err := s.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = row.Exec(subject.SubjectName)
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

func (s SubjectRepoImpl) UpdateSubject(id int, subject *models.Subject) error {
	data, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	query := "UPDATE subject SET subject_name = ? WHERE id = ?"
	row, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = row.Exec(subject.SubjectName, id)
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

func (s SubjectRepoImpl) DeleteSubject(id int) error {
	data, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	query := "DELETE FROM subject WHERE id = ?"
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

func InitSubjectRepoImpl(db *sql.DB) SubjectRepository {
	return &SubjectRepoImpl{db}
}
