package file

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"lab/exp1/src/controllers/home"

	"github.com/gofiber/fiber/v2"
	"github.com/harrytruong/gtfs-sqlite/gtfsconv"
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

	var allFiles []string

	file, err := c.FormFile("fichier")
	if err != nil {
		panic(err)
	}

	zipFile, err := file.Open()
	if err != nil {
		return err
	}

	// Lire le contenu du fichier ZIP en mémoire
	zipData, err := ioutil.ReadAll(zipFile)
	if err != nil {
		return err
	}

	// Créer un lecteur de fichier ZIP à partir des données en mémoire
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		rc, _ := file.Open()
		scanner := bufio.NewScanner(rc)

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}

		defer rc.Close()

		allFiles = append(allFiles, file.Name)
	}

	// if err := c.SaveFile(file, path.Join("uploads", file.Filename)); err != nil {
	// 	fmt.Println(err)
	// 	return c.JSON(fiber.Map{"message": err.Error()})
	// }
	return c.JSON(fiber.Map{"message": "FIle uploaded", "file": allFiles})
}

// processGFTS godoc
// @Summary Process GTFS File
// @Description Route which permit to process a gtfs file that you have uploaded
// @Tags file
// @Router /processGtfs [get]
func processGFTS(c *fiber.Ctx) error {
	gtfsconv.GoBuild(gtfsconv.Options{GTFS: "uploads/abidjan.zip", Dir: "results"})
	return c.JSON(home.HomeRes{Message: "File processing"})
}
