package db

import (
	"database/sql"
	"fmt"
	"time"
	"training-api/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	fmt.Println(connectionString)
	db, err = sql.Open("mysql", connectionString)
	db.SetConnMaxLifetime(time.Minute * 4)
	if err != nil {
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
	fmt.Println("no error encountered")
}

func CreateCon() *sql.DB {
	return db
}
