package customerDB

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateDB() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}
	userName := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	schema := os.Getenv("DATABASE")
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", userName, password, host, schema)
	client, sErr:= sql.Open("mysql", datasource)
	if sErr != nil {
		panic(sErr)
	}
	if pErr := client.Ping(); sErr != nil {
		panic(pErr)

	}

	return client
}
