package migration

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"sync"
)

var onlyOnce sync.Once

func Migrate(connection *gorm.DB) {

	onlyOnce.Do(func() {

		log.Info("Migrating the database...")

		err := connection.AutoMigrate()
		if err != nil {
			log.Error("Error migrating the database: ", err)
		} else {
			log.Info("Database migration is successful.")
		}

	})

}
