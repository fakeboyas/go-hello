package main

import "github.com/gofiber/fiber/v2"

func HelloWorld(res *fiber.Ctx) error {
	return res.SendString("Hello World!")
}

func main() {
	app := fiber.New()

	app.Get("/", HelloWorld)

	app.Listen(":4000")
}
