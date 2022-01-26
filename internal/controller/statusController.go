package controller

import (
	"ExerciseTasks/internal/models"
	r "ExerciseTasks/internal/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStatus(c echo.Context) error {
	s := r.NewStatusRepository()
	status, _ := s.ListStatus()
	var response = models.ResponseStatus{true, "", status, make([]string, 0)}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("json %v\n", string(jsonData))
	return c.JSON(http.StatusOK, string(jsonData))
}
