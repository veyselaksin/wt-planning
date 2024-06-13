package baserouter

import (
	"wt-planning/cmd/basecontroller"
	"wt-planning/db/connection"
	messageRouter "wt-planning/pkg/message/router"

	"github.com/gofiber/fiber/v2"
)

func InitializeRouters(app fiber.Router, client connection.Client) {

	// API v1 routes
	api := app.Group("/v1")

	// Health check routes
	basecontroller.HealthCheck(api)
	basecontroller.SwaggerController(api)

	messageRouter.InitializeRouter(api, client)

}
