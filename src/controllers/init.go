package controllers

import (
	"lab/exp1/src/controllers/file"
	"lab/exp1/src/controllers/home"
	"lab/exp1/src/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendStatus(302)
		return c.Redirect("/api/v1")
	})

	api := app.Group("/api").Group("/v1")

	home.Routes(api)
	user.Routes(api)
	file.Routes(api)

}
