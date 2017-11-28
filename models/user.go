package models

import (
	"errors"
	"go-echo-vue/config"
	"log"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	jwt "github.com/dgrijalva/jwt-go"
)

type (
	// User ...
	User struct {
		ID       int       `storm:"id,increment" json:"id"`
		Name     string    `json:"name" validate:"required"`
		Email    string    `json:"email" validate:"required,email"`
		Password string    `json:"password" validate:"required,min=8"`
		Created  time.Time `json:"created"`
		Updated  time.Time `json:"updated"`
	}

	// JwtCustomClaims ...
	JwtCustomClaims struct {
		ID int `json:"id"`
		jwt.StandardClaims
	}

	// Token ...
	Token struct {
		Token string `json:"token"`
	}
)

// CreateUser ...
func CreateUser(param User) (res Token, err error) {
	user := param
	if err := db.One("Email", user.Email, &user); err == storm.ErrNotFound {
		if err := db.Save(&user); err != nil {
			return res, err
		}
		token, err := createToken(user.ID)
		if err != nil {
			return res, err
		}

		res = Token{}
		res.Token = token
		return res, err
	}

	err = errors.New("Email already exist")
	return res, err
}

// Login ...
func Login(param User) (res Token, err error) {
	user := param
	query := db.Select(q.Eq("Email", user.Email), q.Eq("Password", user.Password))
	query.First(&user)

	if user.ID == 0 {
		return res, err
	}

	token, err := createToken(user.ID)
	if err != nil {
		return res, err
	}

	res = Token{}
	res.Token = token
	return res, err
}

// createToken ...
func createToken(id int) (res string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = config.JwtExp
	res, err = token.SignedString([]byte("secret"))
	if err != nil {
		return res, err
	}
	return res, err
}

// SaveUser ...
func SaveUser(params User) (res User, err error) {
	if err = validate.Struct(params); err != nil {
		log.Printf("data : %v", err)
		return res, err
	}

	if err := db.Save(&params); err != nil {
		return res, err
	}
	return params, err
}

// FindUser ...
func FindUser(id int) User {
	user := User{}
	err = db.One("ID", id, &user)
	return user
}
