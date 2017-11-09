package task

import (
	"go-echo-vue/config"
	"log"
	"net/http"
	"strconv"

	"go-echo-vue/models"

	"github.com/labstack/echo"
)

// GetTasks endpoint
func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		query := c.QueryParam("q")
		pages := c.QueryParam("pages")
		if pages == "" {
			pages = "1"
		}

		var number int
		number, _ = strconv.Atoi(pages)
		data := models.SearchTask((int(number)-1)*20, query)

		if size := len(data); size == 0 {
			return c.JSON(http.StatusNotFound, config.NotFound)
		}
		return c.JSON(http.StatusOK, data)
	}
}

func Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data := models.FindTask(int(id))
		if data.ID == 0 {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusNotFound, config.NotFound)
		}
		return c.JSON(http.StatusOK, data)
	}
}

// Create ...
func Create() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		params := new(models.Task)
		if err = c.Bind(params); err != nil {
			return c.JSON(http.StatusNotAcceptable, config.NotAcceptable)
		}

		task := models.Task{
			Name: params.Name,
		}
		data, err := models.CreateTask(task)
		if err != nil {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, data)
	}
}

// Update ...
func Update() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		params := new(models.Task)
		if err = c.Bind(params); err != nil {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusNotAcceptable, config.NotAcceptable)
		}

		nowTask := models.FindTask(int(id))
		if nowTask.ID == 0 {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusNotFound, config.NotFound)
		}

		task := models.Task{
			ID:   int(id),
			Name: params.Name,
		}
		data, err := models.SaveTask(task)
		if err != nil {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, data)
	}
}

func Delete() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		nowTask := models.FindTask(int(id))
		if nowTask.ID == 0 {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusNotFound, config.NotFound)
		}

		data, err := models.DeleteTask(nowTask)
		if err != nil {
			log.Printf("data : %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, data)
	}
}
