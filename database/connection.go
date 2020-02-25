package database

import (
	_"os"
	"database/sql"
	"smart-hr/library/logger"
	_ "github.com/lib/pq"
	"fmt"
	// "hr/config"
)

var DB *sql.DB

func init() {
	//run this only on debug mode
	db, err :=sql.Open("postgres", "postgres://hxqcdernjwgeyd:9317d41965ac95988490a8d0dc90e32e967aeebbf7e5a0eb4517d29c05a06c81@ec2-3-91-112-166.compute-1.amazonaws.com:5432/d2kaa3ithdvke8") 
	// db, err := sql.Open("postgres",os.Getenv("HEROKU_POSTGRESQL_RED_URL"))
	if err != nil {
		fmt.Println(err)
	} 
	DB = db
	logger.Log.Println("Connection to db successfully")
}