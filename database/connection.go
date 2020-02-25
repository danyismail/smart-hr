package database

import (
	"os"
	"database/sql"
	"smart-hr/library/logger"
	_ "github.com/lib/pq"
	"fmt"
	// "hr/config"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("postgres",os.Getenv("HEROKU_POSTGRESQL_RED_URL"))
	if err != nil {
		fmt.Println(err)
	} 
	DB = db
	logger.Log.Println("Connection to db successfully")
}