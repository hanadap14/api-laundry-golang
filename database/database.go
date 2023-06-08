package database

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/laundry"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db = db
}
