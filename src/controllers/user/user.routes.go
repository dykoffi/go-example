package user

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World dady edy!")
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		// configs.Db().Create()
		return c.SendString("In user")
	})
}
