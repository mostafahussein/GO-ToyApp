package main

import (
	"go-echo-vue/config"
	"go-echo-vue/controllers/task"
	"go-echo-vue/controllers/user"
	"go-echo-vue/models"
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

	e.POST("/v1/register", user.Register())
	e.POST("/v1/login", user.Login())

	r := e.Group("/v1/")
	jwtconfig := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(jwtconfig))
	r.GET("user", user.Get())
	r.PUT("user", user.Update())

	r.GET("tasks", task.Index())
	r.GET("tasks/:id", task.Show())
	r.POST("tasks", task.Create())
	r.PUT("tasks/:id", task.Update())
	r.DELETE("tasks/:id", task.Delete())

	log.Fatal(e.Start(config.HOST + config.PORT))
}
