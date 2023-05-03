package file

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	app.Post("/file/upload", updloadFile)
	app.Post("/file/processGtfs", updloadFile)
}
