package controller

import (
	"ExerciseTasks/internal/models"
	r "ExerciseTasks/internal/repository"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStatus(c echo.Context) error {
	var response = models.ResponseStatusArray{}
	s := r.NewStatusRepository()
	status, _ := s.ListStatus()
	response.Success = true
	response.Message = ""
	response.Data = &status
	jsonData, _ := json.Marshal(response)
	return c.JSON(http.StatusOK, string(jsonData))
}
