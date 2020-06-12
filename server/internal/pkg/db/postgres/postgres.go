package database

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("postgres", "postgres:abc@10.0.0.34:5432/hamsterapps")
	if err != nil {
		log.Panic(err)
	}

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

	driver, _ := postgres.WithInstance(DB, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/postgres",
		"postgres",
		driver,
	)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
