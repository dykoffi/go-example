package user

import "github.com/gofiber/fiber/v2"

func Routes(api fiber.Router) {
	api.Post("/login", loginUser)
}
