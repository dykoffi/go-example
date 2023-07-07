package home

import (
	"fmt"
	"lab/exp1/src/configs"
	"lab/exp1/src/model"

	"github.com/gofiber/fiber/v2"
)

type HomeRes struct {
	Message string `json:"message" validate:"required"`
}

// getApiInfo godoc
//
//	@Summary	Get APi Information
//	@Tags		home
//	@Produce	json
//	@Success	200	{object}	HomeRes
//	@Router		/ [get]
func getApiInfo(c *fiber.Ctx) error {
	return c.JSON(HomeRes{Message: "It works"})
}

func getDocs(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.SendFile("public/rapidoc.html")
}

func addUser(c *fiber.Ctx) error {
	db := configs.Db().Create(&model.User{Mail: "koffiedy@gmail.com", Age: 16})

	if err := db.Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"err": err.Error()})
	}

	return c.JSON(map[string]string{"message": "User added successfully"})
}

func listUsers(ctx *fiber.Ctx) error {
	listUser := configs.Db().First(&model.User{})

	fmt.Println(listUser)

	if err := listUser.Error; err != nil {
		return ctx.Status(400).JSON(map[string]string{"err": err.Error()})
	}

	return ctx.JSON(listUser)

}

func createCookie(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"message": "Cookie creation"})
}

func getCookie(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"message": ctx.Cookies("test")})
}

func clearCookie(ctx *fiber.Ctx) error {
	ctx.ClearCookie()
	return ctx.JSON(map[string]string{"message": "Cookie cleaning"})
}
