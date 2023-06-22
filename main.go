package main

import (
	"go-oauth2-google/router"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	router.Init(app)

	app.Listen(":" + os.Getenv("PORT"))
}
