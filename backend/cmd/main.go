package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
	"os/signal"
	"wt-planning/cmd/baserouter"
	"wt-planning/cmd/config"
	"wt-planning/db/connection"
	"wt-planning/docs"
	"wt-planning/i18n"
)

// @title Insider Case Study API
// @version 1.0
// @description This is a server for Insider Case Study API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /v1
func main() {
	client := connection.New()

	app := fiber.New(config.FiberConfig)

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05.000Z",
		TimeZone:   "Europe/Istanbul",
	}))

	// Initialize routes
	baserouter.InitializeRouters(app, client)

	//Swagger Info configuration
	docs.SwaggerInfo.Host = fmt.Sprint(os.Getenv("APP_HOST"), ":", os.Getenv("APP_PORT"))

	// Init i18n
	i18n.InitBundle("./i18n/languages")

	// Start listening on port 8000
	go func() {
		if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	GracefulShutdown(app, client)
}

func GracefulShutdown(app *fiber.App, client connection.Client) {

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Info("Received terminate,graceful shutdown", sig)

	database, err := client.PostgresConnection.DB()
	if err != nil {
		log.Error("PostgreSQL Closing ERROR :", err)
	}

	err = database.Close()
	if err != nil {
		return
	}
	log.Error("PostgreSQL Closed")

	err = app.Shutdown()
	if err != nil {
		return
	}
}
