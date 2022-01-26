package main

import (
	"ExerciseTasks/internal/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// Route / handler function
	e.GET("/tasks", controller.GetTasks)
	e.POST("/tasks", controller.NewTask)
	e.PUT("/tasks/", controller.UpdateTask)
	e.DELETE("/tasks/:id", controller.DeleteTask)

	e.GET("/status", controller.GetStatus)
	e.Logger.Fatal(e.Start(":3000"))
}
