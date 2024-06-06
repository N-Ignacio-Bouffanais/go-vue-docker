package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Client struct {
	baseURL string
}

type ApiResponse struct {
	Message string `json:"message"`
}

func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

func ServerGo() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hola mundo"})
	})

	log.Fatal(app.Listen(":8083"))

}
