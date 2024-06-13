package connection

import (
	"fmt"
	"os"
	"wt-planning/db/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s application_name='%s' sslmode=disable timezone=Europe/Istanbul",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_USERNAME"),
		os.Getenv("POSTGRES_DB_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_DB_APP_NAME"),
	)

	connection, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return nil
	}

	migration.Migrate(connection)

	return connection
}
