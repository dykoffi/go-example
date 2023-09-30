package controllers

import (
	"lab/exp1/src/controllers/home"
	"lab/exp1/src/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {

	api := app.Group("/api").Group("/v1")

	home.Routes(api)
	user.Routes(api)

	// For UI services
	// app.Static("/", "dist")
	// app.Use(cache.New())
	// app.Get("/*", func(c *fiber.Ctx) error {
	// 	return c.SendFile("dist/index.html")
	// })

}
