package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DBCon *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic("failed to connect to database")
	}

	fmt.Println("Connection to database established")
	db.LogMode(true)
	DBCon = db
	return db
}

func Close(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		panic("cannot close database connection")
	}
	fmt.Println("Connection to database closed")
}

func GetDB() *gorm.DB {
	return DBCon
}
