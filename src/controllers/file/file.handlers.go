package file

import (
	"fmt"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
)

// uploadFile godoc
// @Tags		file
// @Summary	Upload file
// @Accept mpfd
// @Produce json
// @Success	200	{object}	home.HomeRes
// @Param fichier formData file true "File to upload"
// @Router /file/upload [post]
func updloadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("fichier")
	if err != nil {
		panic(err)
	}
	cwd, _ := os.Getwd()
	fmt.Println(c.MultipartForm())

	c.SaveFile(file, path.Join(cwd, "uploads", file.Filename))
	return c.JSON(fiber.Map{"message": "FIle uploaded"})
}
