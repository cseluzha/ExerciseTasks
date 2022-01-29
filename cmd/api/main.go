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
	e.POST("/tasks/findById", controller.FindTaskById)
	e.POST("/tasks/findByTitle", controller.FindTaskByTitle)

	e.GET("/status", controller.GetStatus)
	e.Logger.Fatal(e.Start(":3000"))
}
