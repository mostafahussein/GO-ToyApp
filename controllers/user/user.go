// // Package user controller
package user

// import (
// 	"encoding/hex"
// 	"errors"
// 	"go-echo-vue/config"
// 	"go-echo-vue/models"
// 	"log"
// 	"net/http"
// 	"time"
// 	"unicode/utf8"

// 	"golang.org/x/crypto/scrypt"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo"
// 	validator "gopkg.in/go-playground/validator.v9"
// )

// func toHashFromScrypt(password string) string {
// 	salt := []byte(config.HashSalt)
// 	converted, _ := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
// 	return hex.EncodeToString(converted[:])
// }

// // Register new user.
// func Register() echo.HandlerFunc {
// 	return func(c echo.Context) (err error) {
// 		user := new(models.User)
// 		if err = c.Bind(user); err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}

// 		validate := validator.New()
// 		if err := validate.Struct(user); err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}

// 		save := models.User{
// 			ID:       user.ID,
// 			Name:     user.Name,
// 			Email:    models.NewNullString(user.Email.String),
// 			Password: models.NewNullString(user.Password.String),
// 		}

// 		save, err = validation(save)
// 		if err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}

// 		data, err := models.CreateUser(save)
// 		if err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}
// 		return c.JSON(http.StatusCreated, data)
// 	}
// }

// func validation(user models.User) (res models.User, err error) {
// 	res = user
// 	validate := validator.New()
// 	if err = validate.Struct(res); err != nil {
// 		return res, err
// 	}

// 	password := res.Password
// 	if !password.Valid || !res.Email.Valid {
// 		err = errors.New("password and email are required,")
// 		return res, err
// 	} else if utf8.RuneCountInString(password.String) < 8 {
// 		err = errors.New("password is short,")
// 		return res, err
// 	}

// 	res.Password.String = toHashFromScrypt(password.String)
// 	return res, err
// }

// /*
// Login method.
//     retun jwt token
// */
// func Login() echo.HandlerFunc {
// 	return func(c echo.Context) (err error) {
// 		param := new(models.User)
// 		if err = c.Bind(param); err != nil {
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}

// 		password := param.Password
// 		if password.Valid {
// 			password.String = toHashFromScrypt(password.String)
// 		}

// 		user := models.User{
// 			Email:    param.Email,
// 			Password: password,
// 		}
// 		data, err := models.Login(user)
// 		if err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}
// 		if data.Token == "" {
// 			log.Printf("data : %v", "error authorize")
// 			return c.JSON(http.StatusUnauthorized, config.Unauthorized)
// 		}
// 		return c.JSON(http.StatusOK, data)
// 	}
// }

// // GetUserInfo return users info from jwt token
// func GetUserInfo(c echo.Context) (res models.UserJSON, err error) {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*models.JwtCustomClaims)

// 	res = models.FindUser(claims.ID)
// 	if res.ID == 0 {
// 		return res, err
// 	}
// 	return res, err
// }

// // Get user info
// func Get() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		data, err := GetUserInfo(c)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}
// 		return c.JSON(http.StatusOK, data)
// 	}
// }

// // Update @method put update user info
// func Update() echo.HandlerFunc {
// 	return func(c echo.Context) (err error) {
// 		posts := new(models.User)
// 		if err = c.Bind(posts); err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusNotAcceptable, config.NotAcceptable)
// 		}

// 		user, err := GetUserInfo(c)
// 		if err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}

// 		save := models.User{
// 			ID:       user.ID,
// 			Name:     posts.Name,
// 			Email:    models.NewNullString(posts.Email.String),
// 			Password: models.NewNullString(posts.Password.String),
// 			Created:  user.Created,
// 			Updated:  time.Now(),
// 		}

// 		save, err = validation(save)
// 		if err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}

// 		data, err := models.SaveUser(save)
// 		if err != nil {
// 			log.Printf("data : %v", err)
// 			return c.JSON(http.StatusBadRequest, config.BadRequest)
// 		}
// 		return c.JSON(http.StatusCreated, data)
// 	}
// }
