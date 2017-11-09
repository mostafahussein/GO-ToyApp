package main

import (
	"go-echo-vue/config"
	"go-echo-vue/controllers/task"
	"log"

	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	logPath := config.ProjectPath + "logs/echo.log"
	fp, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fp,
	}))
	e.File("/", "public/index.html")
	e.GET("/tasks", task.GetTasks())
	e.GET("/tasks/:id", task.Get())
	e.POST("/tasks", task.Create())
	e.PUT("/tasks/:id", task.Update())
	e.DELETE("/tasks/:id", task.Delete())

	log.Fatal(e.Start(config.HOST + config.PORT))
}
