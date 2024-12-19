package config

/*
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// Change these to match your PostgreSQL connection details
	connStr := "postgresql://postgres:postgres@localhost:5432/chat-bot?sslmode=disable"

	// Create a new PostgreSQL DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not open DB connection: %v", err)
	}
	defer db.Close()

	// Initialize the driver for PostgreSQL
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not initialize PostgreSQL driver: %v", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"config/migration/",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf(err.Error())
		log.Fatalf("could not create migration instance: %v", err)
	}

	// Parse command-line arguments to determine action
	if len(os.Args) < 2 {
		log.Fatal("missing migration action (up or down)")
	}

	action := os.Args[1]
	switch action {
	case "up":
		// Apply all available migrations
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("could not apply migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully!")
	case "down":
		// Rollback one migration
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("could not rollback migration: %v", err)
		}
		fmt.Println("Migration rolled back successfully!")
	default:
		log.Fatalf("invalid migration action: %s (expected 'up' or 'down')", action)
	}
}
*/
