package home

import (
	"lab/exp1/src/auth"

	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	api.Get("/", getApiInfo)
	api.Get("/docs", getDocs)
	api.Post("/user", addUser)
	api.Get("/getUser", listUsers)

	//cookies
	api.Get("/testPermission", auth.Protect(auth.HasRoleCreateCookies, auth.PolicyOpts{Logic: true}), createCookie)
}
