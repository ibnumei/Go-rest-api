package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConn() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/course_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("can't connect db")
	}
	return db
}