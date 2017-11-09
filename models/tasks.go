package models

import "log"

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SearchTask ...
func SearchTask(pages int, search string) []Task {
	tasks := []Task{}
	db.Limit(20).Offset(pages).Find(&tasks)
	return tasks
}

// FindTask ...
func FindTask(id int) Task {
	task := Task{}
	db.Find(&task, id)
	return task
}

// CreateTask ...
func CreateTask(params Task) (res Task, err error) {
	if err = validate.Struct(params); err != nil {
		log.Printf("data : %v", err)
		return res, err
	}

	if err := db.Create(&params).Error; err != nil {
		return res, err
	}
	return params, err
}

// SaveTask ...
func SaveTask(params Task) (res Task, err error) {
	if err = validate.Struct(params); err != nil {
		log.Printf("data : %v", err)
		return res, err
	}

	if err := db.Save(&params).Error; err != nil {
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

	if err := db.Delete(&params).Error; err != nil {
		return res, err
	}
	return params, err
}
