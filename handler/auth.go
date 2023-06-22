package handler

import (
	"context"
	"encoding/json"
	"go-oauth2-google/config"
	"go-oauth2-google/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"golang.org/x/oauth2"
)

// Auth fiber handler
func Auth(c *fiber.Ctx) error {
	url := config.Get().AuthCodeURL(utils.UUID(), oauth2.AccessTypeOffline)

	return c.Redirect(url)
}

// Callback to receive google's response
func Callback(c *fiber.Ctx) error {
	token, error := config.Get().Exchange(context.Background(), c.Query("code"))
	if error != nil {
		panic(error)
	}

	client := config.Get().Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	googleResponse := new(model.GoogleResponse)

	if err := json.NewDecoder(resp.Body).Decode(&googleResponse); err != nil {
		log.Println(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"email": googleResponse.Email,
		"name":  googleResponse.Name,
		"login": true,
	})
}
