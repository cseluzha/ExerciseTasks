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

	var response = models.ResponseTask{true, "", nil, make([]string, 0)}

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
	jsonData, err := json.Marshal(response)
	return c.JSON(http.StatusNotModified, string(jsonData))
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
	var response = models.ResponseTasks{true, "Task updated", nil, make([]string, 0)}
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
		var response = models.ResponseTasks{true, "Task deleted", nil, make([]string, 0)}
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
	var response = models.ResponseTasks{true, "Tasks", &tasks, make([]string, 0)}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("json %v\n", string(jsonData))
	return c.JSON(http.StatusOK, string(jsonData))
}

func FindTaskById(c echo.Context) error {
	task := models.Task{}	
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil || len(task.Id) == 0{
		errors := [1]string{err.Error()}
		var response = models.ResponseTask{false, "Invalid parameters, IdTask should not be empty", nil, errors[:]}
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusBadRequest, string(jsonData))
	}

	tr := repository.NewTaskRepository()		
	task, et := tr.FindTaskByID(task.Id.String())	
	if et != nil {
		errors := [1]string{et.Error()}
		var response = models.ResponseTask{false, "Task no found", nil, errors[:]}
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusNotFound, string(jsonData))
	}

	var response = models.ResponseTask{true, "Task found", &task, make([]string, 0)}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}	
	return c.JSON(http.StatusOK, string(jsonData))
}

func FindTaskByTitle(c echo.Context) error {
	task := models.Task{}	
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	log.Printf("Title %v\n", task.Title)
	if err != nil || len(task.Title) == 0 {
		errors := [1]string{err.Error()}
		var response = models.ResponseTasks{false, "Invalid parameters, Title should not be empty", nil, errors[:]}
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusBadRequest, string(jsonData))
	}
	tr := repository.NewTaskRepository()
	tasks, _ := tr.FindTaskByTitle(task.Title)
	if len(tasks) > 1 {
		var response = models.ResponseTasks{true, "Tasks found with that title", &tasks, make([]string, 0)}
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusOK, string(jsonData))
	}
	var response = models.ResponseTasks{true, "No tasks found with that title", nil, make([]string, 0)}
	jsonData, _ := json.Marshal(response)
	return c.JSON(http.StatusNotFound, string(jsonData))
}
