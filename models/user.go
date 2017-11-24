package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type (
	// User ...
	User struct {
		ID       int       `db:"id" json:"id"`
		Name     string    `db:"name" json:"name" gorm:"not null" validate:"required"`
		Email    string    `db:"email" json:"email" gorm:"unique_index" validate:"email"`
		Password string    `db:"password" json:"password" validate:"min=8"`
		Created  time.Time `json:"created" sql:"DEFAULT:current_timestamp"`
		Updated  time.Time `json:"updated" sql:"DEFAULT:current_timestamp"`
	}

	// UserJSON ...
	UserJSON struct {
		ID      int       `json:"id"`
		Name    string    `json:"name"`
		Email   string    `db:"email" json:"email" gorm:"unique_index" validate:"email"`
		Created time.Time `json:"created"`
		Updated time.Time `json:"updated"`
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

// TableName ...
// func (UserJSON) TableName() string {
// 	return "user"
// }

// // CreateUser ...
// func CreateUser(param User) (res Token, err error) {
// 	user := param
// 	if err := db.Create(&user).Error; err != nil {
// 		log.Printf("data : %v", err)
// 		fmt.Println(err)
// 		return res, err
// 	}

// 	token, err := createToken(user.ID)
// 	if err != nil {
// 		return res, err
// 	}

// 	res = Token{}
// 	res.Token = token
// 	return res, err
// }

// // Login ...
// func Login(param User) (res Token, err error) {
// 	user := param
// 	db.Where(&user).First(&user)

// 	if user.ID == 0 {
// 		return res, err
// 	}

// 	token, err := createToken(user.ID)
// 	if err != nil {
// 		return res, err
// 	}

// 	res = Token{}
// 	res.Token = token
// 	return res, err
// }

// // createToken ...
// func createToken(id int) (res string, err error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["id"] = id
// 	claims["exp"] = config.JwtExp
// 	res, err = token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, err
// }

// // SaveUser ...
// func SaveUser(params User) (res User, err error) {
// 	if err = validate.Struct(params); err != nil {
// 		log.Printf("data : %v", err)
// 		return res, err
// 	}

// 	if err := db.Save(&params).Error; err != nil {
// 		return res, err
// 	}
// 	return params, err
// }

// // FindUser ...
// func FindUser(id int) UserJSON {
// 	user := UserJSON{}
// 	db.First(&user, id)
// 	return user
// }
