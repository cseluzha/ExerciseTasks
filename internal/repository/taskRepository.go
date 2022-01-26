package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
	m "ExerciseTasks/internal/models"
)

type taskRepository struct {
	db *sql.DB
}

type TasksRepository interface {
	NewTask(task m.Task) string
	UpdateTask(task m.Task) int64
	DeleteTask(taskId string) int64
	ListTasks() ([]m.Task, error)
}

func (tr *taskRepository) NewTask(task m.Task) string {
	// close database
	defer tr.db.Close()
	insertStmt := `INSERT INTO  practices.Tasks (Id, Title, Description, CreatedOn) VALUES ($1, $2, $3, $4) RETURNING Id`
	var id uuid.UUID

	// Scan function will save the insert id in the id
	err := tr.db.QueryRow(insertStmt, GenerateUUID(), task.Title, task.Description, time.Now()).Scan(&id)
	CheckError(err)
	fmt.Printf("Inserted new task id %v\n", id)
	return id.String()
}

func (tr *taskRepository) UpdateTask(task m.Task) int64 {
	// close database
	defer tr.db.Close()

	// create the update sql query
	updateStmt := `UPDATE  practices.Tasks SET Title=$2, Description=$3, UpdatedOn=$4  WHERE Id=$1`

	// execute the sql statement
	res, err := tr.db.Exec(updateStmt, task.Id, task.Title, task.Description, time.Now())
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

func (tr *taskRepository) DeleteTask(taskId string) int64 {
	// close database
	defer tr.db.Close()

	// create the delete sql query
	deleteStmt := `DELETE FROM practices.Tasks WHERE Id=$1`
	// execute the sql statement
	res, err := tr.db.Exec(deleteStmt, taskId)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func (tr *taskRepository) ListTasks() ([]m.Task, error) {
	// close database
	defer tr.db.Close()

	var tasks []m.Task

	// create the select sql query
	sqlStatement := `SELECT * FROM practices.Tasks WHERE Active=true`
	// execute the sql statement
	rows, err := tr.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var task m.Task

		// unmarshal the row object to user
		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedOn, &task.UpdatedOn)

		CheckError(err)
		// append the user in the users slice
		tasks = append(tasks, task)
	}
	// return empty users on error
	return tasks, err
}

func NewTaskRepository() TasksRepository {
	return &taskRepository{db: CreateConnection()}
}
