package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var onlyOnce sync.Once

func Open() *gorm.DB {
	// Open the database connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_USERNAME"),
		os.Getenv("POSTGRES_DB_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_DB_SSLMODE"),
		os.Getenv("POSTGRES_DB_TIMEZONE"),
	)

	// Postgres connection
	connection, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return nil
	}

	// Migrate the database
	onlyOnce.Do(func() {

		fmt.Println("Migrating the database...")

		err := connection.AutoMigrate()
		if err != nil {
			fmt.Println("Error migrating the database: ", err)
		} else {
			fmt.Println("Database migration is successful.")
		}

	})

	return connection
}
