package middlewares

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	clientId     = os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm        = os.Getenv("KEYCLOAK_REALM")
	hostname     = os.Getenv("KEYCLOAK_HOST")
)

func Keycloak(ctx *fiber.Ctx) error {
	fmt.Println(ctx.ClientHelloInfo())
	return ctx.Next()
}
