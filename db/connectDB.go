package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbConnection() {

	var (
		host     string = os.Getenv("DB_HOST")
		user     string = os.Getenv("DB_USERNAME")
		password string = os.Getenv("DB_PASSWORD")
		dbname   string = os.Getenv("DATABASE_NAME")
		port     string = os.Getenv("DB_PORT")
	)

	var DSN string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	err := recover()
	db, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("\nERROR DATABASE NOT CONNECTED\n", err)
	}

	if err == nil {
		log.Printf("\nDATABASE CONNECTED SUCESSFULLY\n")
	}
}
