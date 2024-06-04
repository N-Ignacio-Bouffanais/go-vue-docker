package main

import (
	"log"
	"net/http"

	"go-vue-docker/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	token := config.GetToken()

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hola mundo"})
	})
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"Telegram Bot Token: %s": token})
	})

	log.Fatal(app.Listen(":8083"))

	fs := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", fs)

	log.Println("Server started on :8083")

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
