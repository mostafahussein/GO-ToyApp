package models

import (
	"log"

	"github.com/asdine/storm"
)

// Task is a struct containing Task data
type Task struct {
	ID   int    `storm:"id,increment" json:"id"`
	Name string `json:"name"`
}

// SearchTask ...
func SearchTask(pages int, search string) []Task {
	tasks := []Task{}
	db.All(&tasks, storm.Limit(20), storm.Skip(pages))
	return tasks
}

// FindTask ...
func FindTask(id int) Task {
	task := Task{}
	err = db.One("ID", id, &task)
	return task
}

// CreateTask ...
func CreateTask(params Task) (res Task, err error) {
	if err = validate.Struct(params); err != nil {
		log.Printf("data : %v", err)
		return res, err
	}
	if err := db.One("Name", params.Name, &params); err == storm.ErrNotFound {
		if err := db.Save(&params); err != nil {
			return res, err
		}
	}
	return params, err

}

// SaveTask ...
func SaveTask(params Task) (res Task, err error) {
	if err = validate.Struct(params); err != nil {
		log.Printf("data : %v", err)
		return res, err
	}

	if err := db.Update(&params); err != nil {
		return res, err
	}
	return params, err
}

// DeleteTask ...
func DeleteTask(params Task) (res Task, err error) {
	if err = validate.Struct(params); err != nil {
		log.Printf("data : %v", err)
		return res, err
	}

	if err := db.DeleteStruct(&params); err != nil {
		return res, err
	}
	return params, err
}
