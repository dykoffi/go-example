package home

import (
	"lab/exp1/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	api.Get("/", getApiInfo)
	api.Get("/docs", getDocs)
	api.Get("/addUser", middlewares.Keycloak, addUser)
	api.Get("/getUser", listUsers)
}
