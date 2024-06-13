package basecontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to Insider Case Study API")
}

// HealthCheck godoc
// @Summary Welcome API
// @Description Welcome to Trendbox POS API
// @Tags Health Check
// @Accept */*
// @Produce application/json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func healthcheck(ctx *fiber.Ctx) error {
	return ctx.SendString("Insider Case Study API is running")
}

func HealthCheck(app fiber.Router) {
	app.Get("/", welcome)
	app.Get("/health-check", healthcheck)
}

func SwaggerController(app fiber.Router) {
	app.Get("/docs/*", swagger.HandlerDefault)
}
