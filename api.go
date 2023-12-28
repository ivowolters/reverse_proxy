package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	fmt.Println("dsafasdf")
	app := fiber.New()

	app.Use(cors.New())

	// app.Use(func(c *fiber.Ctx) error {
	// 	if c.Is("json") {
	// 		return c.Next()
	// 	}
	// 	return c.SendString("Only JSON allowed!")
	// })

	app.Get("/", func(c *fiber.Ctx) error {

		url := fetchUrl()

		proxy.DoRedirects(c, url, 3)

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
