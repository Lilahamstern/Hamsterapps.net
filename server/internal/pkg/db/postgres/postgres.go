package database

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("postgres", os.Getenv("HAMSTERAPPS_DB_URL"))
	if err != nil {
		log.Panic(err)
	}
	//defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	DB = db
}

func PingCheckDB() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	PingCheckDB()
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/postgres",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
