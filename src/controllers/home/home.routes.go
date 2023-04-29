package home

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	api.Get("/", getApiInfo)
	api.Get("/docs", getDocs)
}
