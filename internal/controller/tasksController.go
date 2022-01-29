package controller

import (
	models "ExerciseTasks/internal/models"
	repository "ExerciseTasks/internal/repository"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func NewTask(c echo.Context) error {
	task := models.Task{}
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	defer c.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	var response = models.ResponseTask{}
	tr := repository.NewTaskRepository()
	task = models.Task{
		Id:          repository.GenerateUUID(),
		Title:       task.Title,
		Description: task.Description,
	}
	id := tr.NewTask(task)
	if len(id) > 0 {
		response.Data = &task

		response.Message = "task created successfully"
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Failed reading the request body %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		}
		log.Printf("json %#v\n", string(jsonData))
		return c.JSON(http.StatusOK, string(jsonData))
	}
	response.Success = false
	response.Message = "Task not created"
	jsonData, _ := json.Marshal(response)
	return c.JSON(http.StatusNotModified, string(jsonData))
}

func UpdateTask(c echo.Context) error {
	task := models.Task{}
	var response = models.ResponseTask{}
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
	if val >= 0 {
		response.Success = true
		response.Message = "Task updated"
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusOK, string(jsonData))
	}
	response.Success = false
	response.Message = "Task cann't updated"
	jsonData, _ := json.Marshal(response)
	return c.JSON(http.StatusNotModified, string(jsonData))
}

func DeleteTask(c echo.Context) error {
	var response = models.ResponseTasks{}
	idtask := c.QueryParam("id")
	dataType := c.Param("data")
	if dataType == "json" {
		tr := repository.NewTaskRepository()
		val := tr.DeleteTask(idtask)
		if val >= 0 {
			response.Success = true
			response.Message = "Task deleted"
			jsonData, _ := json.Marshal(response)
			return c.JSON(http.StatusOK, string(jsonData))
		}
		response.Success = false
		response.Message = "Task not deleted"
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusNotModified, string(jsonData))
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data",
		})
	}
}

func GetTasks(c echo.Context) error {
	tr := repository.NewTaskRepository()
	var response = models.ResponseTasks{}
	tasks, _ := tr.ListTasks()
	if len(tasks) > 0 {
		response.Success = true
		response.Data = &tasks
		response.Message = "Tasks"
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusOK, string(jsonData))
	}
	response.Success = false
	response.Message = "Tasks not found"
	jsonData, _ := json.Marshal(response)
	return c.JSON(http.StatusOK, string(jsonData))

}

func FindTaskById(c echo.Context) error {
	task := models.Task{}
	var response = models.ResponseTask{}
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil || len(task.Id) == 0 {
		errors := [1]string{err.Error()}
		response.Success = false
		response.Message = "Invalid parameters, IdTask should not be empty"
		response.Errors = errors[:]
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusBadRequest, string(jsonData))
	}

	tr := repository.NewTaskRepository()
	task, et := tr.FindTaskByID(task.Id.String())
	if et != nil {
		errors := [1]string{et.Error()}
		response.Success = false
		response.Message = "Task no found"
		response.Errors = errors[:]
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusNotFound, string(jsonData))
	}

	response.Success = true
	response.Data = &task
	response.Message = "Task found"
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, string(jsonData))
}

func FindTaskByTitle(c echo.Context) error {
	task := models.Task{}
	var response = models.ResponseTasks{}
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil || len(task.Title) == 0 {
		errors := [1]string{err.Error()}
		response.Success = false
		response.Message = "Invalid parameters, Title should not be empty"
		response.Errors = errors[:]
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusBadRequest, string(jsonData))
	}
	tr := repository.NewTaskRepository()
	tasks, _ := tr.FindTaskByTitle(task.Title)
	if len(tasks) > 0 {
		response.Success = true
		response.Message = "Tasks found with that title"
		response.Data = &tasks
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusOK, string(jsonData))
	}
	response.Success = true
	response.Message = "No tasks found with that title"
	jsonData, _ := json.Marshal(response)
	return c.JSON(http.StatusNotFound, string(jsonData))
}
