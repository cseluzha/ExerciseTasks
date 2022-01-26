package controller

import (
	models "ExerciseTasks/internal/models"
	repository "ExerciseTasks/internal/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func NewTask(c echo.Context) error {
	task := models.Task{}
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	defer c.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	var response = models.ResponseTasks{true, "", make([]models.Task, 1), make([]string, 0)}

	tr := repository.NewTaskRepository()
	id := tr.NewTask(models.Task{
		Id:          repository.GenerateUUID(),
		Title:       task.Title,
		Description: task.Description,
	})
	response.Data[0].Id = uuid.MustParse(id)
	response.Data[0].Title = task.Title
	response.Data[0].Description = task.Description

	response.Message = "task created successfully"
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("json %#v\n", string(jsonData))
	return c.JSON(http.StatusOK, jsonData)
}

func UpdateTask(c echo.Context) error {
	task := models.Task{}
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	tr := repository.NewTaskRepository()
	val := tr.UpdateTask(models.Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	})
	var response = models.ResponseTasks{true, "Task updated", make([]models.Task, 0), make([]string, 0)}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("json %v\n %v", string(jsonData), val)
	return c.JSON(http.StatusOK, string(jsonData))
}

func DeleteTask(c echo.Context) error {
	idtask := c.QueryParam("id")
	dataType := c.Param("data")
	if dataType == "json" {
		tr := repository.NewTaskRepository()
		val := tr.DeleteTask(idtask)
		var response = models.ResponseTasks{true, "Task deleted", make([]models.Task, 0), make([]string, 0)}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Failed reading the request body %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		}
		log.Printf("json %v\n %v", string(jsonData), val)
		return c.JSON(http.StatusOK, string(jsonData))
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data",
		})
	}
}

func GetTasks(c echo.Context) error {
	tr := repository.NewTaskRepository()
	tasks, _ := tr.ListTasks()
	var response = models.ResponseTasks{true, "Tasks", tasks, make([]string, 0)}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("json %v\n", string(jsonData))
	return c.JSON(http.StatusOK, string(jsonData))
}
