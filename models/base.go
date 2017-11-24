// Package models ...
package models

import (
	"github.com/asdine/storm"
	validator "gopkg.in/go-playground/validator.v9"
)

var db, err = storm.Open("my.db")

var validate = validator.New()
