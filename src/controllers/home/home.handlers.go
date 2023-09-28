package home

import (
	"fmt"
	"lab/exp1/src/db"

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

// addUser godoc
//
//		@Summary	Add user in the system
//		@Description	Add the user that you want with all information describe in the minefields
//		@Tags		home
//		@Produce	json
//		@Success	200	{object}	HomeRes
//	 @Param  range formData string true "File to upload"
//	 @Param  name formData string true "For user name"
//	 @Param  age formData number true "For user age"
//	 @Param  file formData file true "File to upload"
//		@Router		/user [post]
func addUser(c *fiber.Ctx) error {
	// fmt.Println(c.FormValue())
	db := db.Db().Create(&db.User{Mail: "koffiedy@gmail.com", Age: 16})

	if err := db.Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"err": err.Error()})
	}

	return c.JSON(map[string]string{"message": "User added successfully"})
}

func listUsers(ctx *fiber.Ctx) error {
	listUser := db.Db().First(&db.User{})

	fmt.Println(listUser)

	if err := listUser.Error; err != nil {
		return ctx.Status(400).JSON(map[string]string{"err": err.Error()})
	}

	return ctx.JSON(listUser)

}

func createCookie(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"message": "Cookie creation"})
}
