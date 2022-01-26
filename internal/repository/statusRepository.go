package repository

import (
	m "ExerciseTasks/internal/models"
	"database/sql"	
)

type statusRepository struct {
	db *sql.DB
}

type StatusRepository interface {
	ListStatus() ([]m.Status, error)
}

func (tr *statusRepository) ListStatus() ([]m.Status, error) {
	// close database
	defer tr.db.Close()

	var status []m.Status

	// create the select sql query
	sqlStatement := `SELECT * FROM practices."StatusTask"`
	// execute the sql statement
	rows, err := tr.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var s m.Status

		// unmarshal the row object to user
		err = rows.Scan(&s.Id, &s.Name)

		CheckError(err)
		// append the user in the users slice
		status = append(status, s)
	}
	// return empty users on error
	return status, err
}

func NewStatusRepository() StatusRepository {
	return &statusRepository{db: CreateConnection()}
}
