package config

import (
	"errors"
	"wt-planning/helpers/genericresponse"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var FiberConfig = fiber.Config{
	AppName:   " Insider Case Study API",
	BodyLimit: 1024 * 1024 * 50, // 50 MB

	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		var code int = fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a fiber.*Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		log.Error("Error occurred: ", err)

		return baseresponse.ErrorResponse(ctx, code, "Unexpected error occurred")
	},
}
