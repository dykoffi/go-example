package home

import (
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
