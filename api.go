package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		url := fetchUrl()
		proxy.DoRedirects(c, url, 3)
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
