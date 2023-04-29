package controllers

import (
	"lab/exp1/src/controllers/file"
	"lab/exp1/src/controllers/home"
	"lab/exp1/src/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {

	api := app.Group("/api").Group("/v1")

	home.Routes(api)
	user.Routes(api)
	file.Routes(api)

}
