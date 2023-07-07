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
	api.Get("/postCookie",
		auth.Protect(
			auth.Policy{
				Roles: []string{},
				Users: []string{"9d9c4590-098e-4645-a487-aa5b2ea76ce7"},
				Times: auth.TimeInterval{StartTime: "2023-07-04 10:43:00", ExpireTime: "2023-07-08 12:00:00"},
			},
			auth.PolicyOpts{DecisionStrategy: "Unanymous", Logic: true}),
		createCookie)
	api.Get("/getCookie", getCookie)
	api.Get("/delCookie", clearCookie)
}
