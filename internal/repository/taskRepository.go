package repository

import (
	m "ExerciseTasks/internal/models"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type taskRepository struct {
	db *sql.DB
}

type TasksRepository interface {
	NewTask(task m.Task) string
	UpdateTask(task m.Task) int64
	DeleteTask(taskId string) int64
	ListTasks() ([]m.Task, error)
	FindTaskByID(taskId string) (m.Task, error)
	FindTaskByTitle(title string) ([]m.Task, error)
}

func (tr *taskRepository) NewTask(task m.Task) string {
	// close database
	defer tr.db.Close()
	insertStmt := `INSERT INTO practices."Tasks" ("Id", "Title", "Description", "CreatedOn") VALUES ($1, $2, $3, $4) RETURNING "Id"`
	var Id uuid.UUID

	// Scan function will save the insert id in the id
	err := tr.db.QueryRow(insertStmt, GenerateUUID(), task.Title, task.Description, time.Now()).Scan(&Id)
	CheckError(err)
	fmt.Printf("Inserted new task id %v\n", Id)
	return Id.String()
}

func (tr *taskRepository) UpdateTask(task m.Task) int64 {
	// close database
	defer tr.db.Close()
	var updateStmt string = ""
	var res sql.Result = nil
	var err error = nil
	var rowsAffected int64 = 0

	if task.Title == "" && task.Description == "" && IsValidUUID(task.Status.Id.String()) {
		updateStmt = `UPDATE "practices"."Tasks" SET  "StatusId"=$1, "UpdatedOn"=$2  WHERE "Id"::text = '` + strings.ReplaceAll(task.Id.String(), "\n", "") + `'`
		// execute the sql statement
		res, err = tr.db.Exec(updateStmt, task.Status.Id, time.Now())
		CheckError(err)
		// check how many rows affected
		rowsAffected, _ := res.RowsAffected()
		fmt.Printf("Total rows/record affected %v", rowsAffected)
		return rowsAffected
	}

	if task.Title == "" && task.Description != "" && !IsValidUUID(task.Status.Id.String()) {
		updateStmt = `UPDATE "practices"."Tasks" SET  "Description"=$1, "UpdatedOn"=$2  WHERE "Id"::text = '` + strings.ReplaceAll(task.Id.String(), "\n", "") + `'`
		// execute the sql statement
		res, err = tr.db.Exec(updateStmt, task.Description, time.Now())
		CheckError(err)
		// check how many rows affected
		rowsAffected, _ := res.RowsAffected()
		fmt.Printf("Total rows/record affected %v", rowsAffected)
		return rowsAffected
	}

	if task.Title == "" && task.Description != "" && IsValidUUID(task.Status.Id.String()) {
		updateStmt = `UPDATE "practices"."Tasks" SET "Description"=$1, "UpdatedOn"=$2, "StatusId"=$3  WHERE "Id"::text = '` + strings.ReplaceAll(task.Id.String(), "\n", "") + `'`
		// execute the sql statement
		res, err = tr.db.Exec(updateStmt, task.Description, time.Now(), task.Status.Id)
		CheckError(err)
		// check how many rows affected
		rowsAffected, _ := res.RowsAffected()
		fmt.Printf("Total rows/record affected %v", rowsAffected)
		return rowsAffected
	}

	if task.Title != "" && task.Description == "" && !IsValidUUID(task.Status.Id.String()) {
		updateStmt = `UPDATE "practices"."Tasks" SET "Title"=$1, "UpdatedOn"=$2  WHERE "Id"::text = '` + strings.ReplaceAll(task.Id.String(), "\n", "") + `'`
		// execute the sql statement
		res, err = tr.db.Exec(updateStmt, task.Title, time.Now())
		CheckError(err)
		// check how many rows affected
		rowsAffected, _ := res.RowsAffected()
		fmt.Printf("Total rows/record affected %v", rowsAffected)
		return rowsAffected
	}

	if task.Title != "" && task.Description == "" && IsValidUUID(task.Status.Id.String()) {
		updateStmt = `UPDATE "practices"."Tasks" SET "Title"=$1, "UpdatedOn"=$2, "StatusId"=$3 WHERE "Id"::text = '` + strings.ReplaceAll(task.Id.String(), "\n", "") + `'`
		// execute the sql statement
		res, err = tr.db.Exec(updateStmt, task.Title, time.Now(), task.Status.Id)
		CheckError(err)
		// check how many rows affected
		rowsAffected, _ := res.RowsAffected()
		fmt.Printf("Total rows/record affected %v", rowsAffected)
		return rowsAffected
	}

	if task.Title != "" && task.Description != "" && !IsValidUUID(task.Status.Id.String()) {
		updateStmt = `UPDATE "practices"."Tasks" SET "Title"=$1, "UpdatedOn"=$2, "Description"=$3 WHERE "Id"::text = '` + strings.ReplaceAll(task.Id.String(), "\n", "") + `'`
		// execute the sql statement
		res, err = tr.db.Exec(updateStmt, task.Title, time.Now(), task.Description)
		CheckError(err)
		// check how many rows affected
		rowsAffected, _ := res.RowsAffected()
		fmt.Printf("Total rows/record affected %v", rowsAffected)
		return rowsAffected
	}

	if task.Title != "" && task.Description != "" && IsValidUUID(task.Status.Id.String()) {
		updateStmt = `UPDATE "practices"."Tasks" SET "Title"=$1, "Description"=$2, "StatusId"=$3, "UpdatedOn"=$4  WHERE "Id"::text = '` + strings.ReplaceAll(task.Id.String(), "\n", "") + `'`
		// execute the sql statement
		res, err = tr.db.Exec(updateStmt, task.Title, task.Description, task.Status.Id, time.Now())
		CheckError(err)
		// check how many rows affected
		rowsAffected, _ := res.RowsAffected()
		fmt.Printf("Total rows/record affected %v", rowsAffected)
		return rowsAffected
	}	
	return rowsAffected
}

func (tr *taskRepository) DeleteTask(taskId string) int64 {
	// close database
	defer tr.db.Close()

	// create the delete sql query
	deleteStmt := `UPDATE "practices"."Tasks" SET "Active"=false, "UpdatedOn"=$1  WHERE "Id"::text = '` + strings.ReplaceAll(taskId, "\n", "") + `'`
	fmt.Printf("deleteStmt: %v \n", deleteStmt)
	// execute the sql statement
	res, err := tr.db.Exec(deleteStmt, time.Now())
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	return rowsAffected
}

func (tr *taskRepository) ListTasks() ([]m.Task, error) {
	// close database
	defer tr.db.Close()

	var tasks []m.Task

	// create the select sql query
	sqlStatement := `SELECT * FROM practices."Tasks" WHERE "Active"=true`
	// execute the sql statement
	rows, err := tr.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()
	sr := NewStatusRepository()
	// iterate over the rows
	for rows.Next() {
		var task m.Task
		var statusId string = ""
		// unmarshal the row object to user
		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedOn, &task.UpdatedOn, &statusId, &task.Active)

		CheckError(err)
		// append the user in the users slice
		task.Status = sr.GetStatus(statusId)
		tasks = append(tasks, task)
	}
	// return empty users on error
	return tasks, err
}

func (tr *taskRepository) FindTaskByID(taskId string) (m.Task, error) {
	// close database
	defer tr.db.Close()
	var task m.Task
	// create the select sql query
	sqlStatement := `SELECT * FROM practices."Tasks" WHERE "Id"=$1 AND "Active"=true`
	// execute the sql statement
	var statusId string = ""
	rows := tr.db.QueryRow(sqlStatement, taskId)
	err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedOn, &task.UpdatedOn, &statusId, &task.Active)
	if err == nil {
		sr := NewStatusRepository()
		task.Status = sr.GetStatus(statusId)
	}
	return task, err
}

func (tr *taskRepository) FindTaskByTitle(title string) ([]m.Task, error) {
	// close database
	defer tr.db.Close()

	var tasks []m.Task
	// create the select sql query
	sqlStatement := `SELECT * FROM practices."Tasks" WHERE  "Title" ILIKE  '%' || $1 || '%' AND "Active"=true;`
	// execute the sql statement
	rows, err := tr.db.Query(sqlStatement, &title)
	CheckError(err)
	// close the statement
	defer rows.Close()
	sr := NewStatusRepository()
	// iterate over the rows
	for rows.Next() {
		var task m.Task
		var statusId string = ""
		// unmarshal the row object to user
		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedOn, &task.UpdatedOn, &statusId, &task.Active)

		CheckError(err)
		// append the user in the users slice
		task.Status = sr.GetStatus(statusId)
		tasks = append(tasks, task)
	}
	// return empty users on error
	return tasks, err
}

func NewTaskRepository() TasksRepository {
	return &taskRepository{db: CreateConnection()}
}
