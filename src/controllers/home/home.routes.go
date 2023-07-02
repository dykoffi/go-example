package home

import (
	"lab/exp1/src/auth"
	"lab/exp1/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	api.Get("/", getApiInfo)
	api.Get("/docs", getDocs)
	api.Get("/addUser", middlewares.Keycloak, addUser)
	api.Get("/getUser", listUsers)

	//cookies
	api.Get("/postCookie", auth.Protect(auth.Policy{Users: []string{}, Roles: []string{}}, "", ""), createCookie)
	api.Get("/getCookie", getCookie)
	api.Get("/delCookie", clearCookie)
}
