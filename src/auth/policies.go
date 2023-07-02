package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TimeInterval = [2]time.Time

type Policy struct {
	Users []string `json:"users"`
	Roles []string `json:"roles"`
	Times string   `json:"times"`
}

func enforcePolicy(token *jwt.MapClaims, policy Policy, decisionStrategy string, logic string) (bool, error) {

	return true, nil
}

func applyTimePolicy(interval TimeInterval) {

}

func applyRolePolicy() {

}

func applyUserPolicy() {

}

func Protect(policy Policy, decisionStrategy string, logic string) func(*fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {

		userToken, err := GetBearerToken(ctx.Get("Authorization"))

		if err != nil {

		}

		_, tokenClaims, err := KC.Gocloak.DecodeAccessToken(ctx.Context(), userToken, KC.Realm)

		if err != nil {

		}

		result, errors := enforcePolicy(tokenClaims, policy, decisionStrategy, logic)

		if errors != nil {

		}

		if result == false {

		}

		return ctx.Next()
	}
}
