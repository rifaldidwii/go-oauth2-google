package router

import (
	"go-oauth2-google/handler"

	"github.com/gofiber/fiber/v2"
)

// Routes for fiber
func Init(app *fiber.App) {
	app.Get("/", handler.Auth)

	app.Get("/callback", handler.Callback)
}
