package controllers

import (
	"lab/exp1/src/controllers/home"

	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {

	// permissions := rptResult.Permissions

	api := app.Group("/api").Group("/v1")

	home.Routes(api)

	// For UI services
	app.Static("/", "dist")
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("dist/index.html")
	})

}
