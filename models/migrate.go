package models

func init() {
	db := DB()
	db.AutoMigrate(&Task{})
}
