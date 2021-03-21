package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("when there is a shit going to move iut of here to")
	})

	app.Get("/:param", func(c *fiber.Ctx) {
		c.Send("pAraMi iS " + c.Params("param"))
	})
// this is a test comment
	app.Listen(8080)
}
