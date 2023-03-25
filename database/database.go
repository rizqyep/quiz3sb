package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	once sync.Once
)

func NewDBConnection() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("DB connection failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}
	return db
}

func GetDBConnection() *sql.DB {
	once.Do(func() {
		DB = NewDBConnection()
	})
	return DB
}
