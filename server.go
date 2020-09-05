package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("No HElLo foR tHe wOrLD")
	})

	app.Get("/:param", func(c *fiber.Ctx) {
		c.Send("pAraMi iS " + c.Params("param"))
	})

	app.Listen(8080)
}
