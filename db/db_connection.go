package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// create mysql connection

func CreateConn() *sql.DB {
	db, err := sql.Open("mysql", "app:password@tcp(localhost:9000)/app")
	if err != nil {
		fmt.Println("Error when connect to MySQL Database ", err.Error())
	}

	fmt.Println("Connected to database localhost")
	err = db.Ping()

	if err != nil {
		fmt.Println("DB is not connected")
		fmt.Println(err.Error())
	}
	return db
}
