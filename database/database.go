package database

import (
	"database/sql"
	"fmt"
	"quiz3-rizqyep/utils"
	"sync"

	"github.com/gobuffalo/packr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DB   *sql.DB
	once sync.Once
)

func NewDBConnection() *sql.DB {
	envMap := utils.InitEnv()
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		envMap["PGHOST"], envMap["PGPORT"], envMap["PGUSER"], envMap["PGPASSWORD"], envMap["PGDATABASE"])

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

func DBMigrate() {
	dbParam := GetDBConnection()
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)

	if errs != nil {
		panic(errs)
	}

	fmt.Println("Applied ", n, " migrations!")
}
