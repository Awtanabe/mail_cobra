package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetUp() {
	name := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("MYSQL_PROT")
	db_name := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", name, pass, host, port, db_name)
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			return
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
