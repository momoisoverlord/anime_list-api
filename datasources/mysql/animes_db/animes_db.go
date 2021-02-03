package animes_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Client *sql.DB
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("mysqlUsersUsername")
	password := os.Getenv("mysqlUsersPassword")
	host := os.Getenv("mysqlUsersHost")
	schema := os.Getenv("mysqlUsersSchema")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		//panic(err)
		return
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured...")
}
