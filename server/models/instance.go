package models

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB;

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbName := os.Getenv("db_name")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	tmp, err := sql.Open("mysql", dbUri);
	db = tmp;
	if err != nil {
		log.Fatal("Couldn't connect to the database.");
	}
}

func GetDB() *sql.DB {
	return db
}
