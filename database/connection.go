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
	// connStr := "user=postgres dbname=smart_hr sslmode=disable"
	db, err := sql.Open("postgres",os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
	} 
	DB = db
	logger.Log.Println("Connection to db successfully")
}