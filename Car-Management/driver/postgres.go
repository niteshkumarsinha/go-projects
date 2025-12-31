package driver

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
				os.Getenv("DB_HOST"), 
				os.Getenv("DB_PORT"), 
				os.Getenv("DB_USER"), 
				os.Getenv("DB_PASSWORD"), 
				os.Getenv("DB_NAME"))

	fmt.Println("Waiting to connect to the database")
	time.Sleep(5 * time.Second)

	var err error
	db, err = sql.Open("postgres", connStr)
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("Error connecting to the database", err)
		}
	}()
	if err != nil {
		panic(err)
	}	
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database")	
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		fmt.Println("Error closing the database connection", err)
	}
}
