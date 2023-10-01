package auth

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx, username string, password string) (*gocloak.JWT, error) {

	return kc.Gocloak.Login(
		ctx.Context(),
		kc.ClientId,
		kc.ClientSecret,
		kc.Realm,
		username,
		password,
	)

}

func Logout(ctx *fiber.Ctx, refreshToken string) error {
	return kc.Gocloak.Logout(
		ctx.Context(),
		kc.ClientId,
		kc.ClientSecret,
		kc.Realm,
		refreshToken,
	)

}

func RefreshToken(ctx *fiber.Ctx, refreshToken string) (*gocloak.JWT, error) {
	return kc.Gocloak.RefreshToken(
		ctx.Context(),
		refreshToken,
		kc.ClientId,
		kc.ClientSecret,
		kc.Realm,
	)
}
