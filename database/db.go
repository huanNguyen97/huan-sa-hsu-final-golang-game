package database

import (
	// Go importing module
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect db
func Db_Connection() *gorm.DB {
	dsn := "b79e58ead43f69:33113388@tcp(us-cdbr-east-04.cleardb.com)/heroku_83cb13da50f38e0"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("?")
	}
	return db
}