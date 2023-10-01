package user

import (
	"lab/exp1/src/auth"

	"github.com/gofiber/fiber/v2"
)

// LoginUSer			godoc
//
//	@Tags								User
//	@Router							/login [post]
//	@Summary						Log user to the system
//	@Description				Pass user mail and password for geting token auth
//	@Produce						json
//	@Param  						mail formData string true "User mail" format(email)
//	@Param  						password formData string true "User password"
//	@Success 						200	{object}	SuccessRes
//	@Failure 						401	{object}	ErrorRes
//	@Failure 						500	{object}	ErrorRes
//	@Security						ApiKeyAuth
func loginUser(ctx *fiber.Ctx) error {

	credentials := new(Creds)

	if err := ctx.BodyParser(credentials); err != nil {
		return err
	}

	userInfo, err := auth.Login(ctx, credentials.Mail, credentials.Password)

	if err != nil {
		return ctx.Status(401).JSON(map[string]string{"error": "InvalidUserError", "message": err.Error()})
	} else {
		return ctx.JSON(userInfo)
	}

}
