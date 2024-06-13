package helpers

import "github.com/gofiber/fiber/v2"

func GetLanguage(ctx *fiber.Ctx) string {
	lang := ctx.Get("Accept-Language")
	if lang == "" {
		lang = "en"
	}
	return lang
}
