package models

func init() {
	db := DB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Task{})
}
