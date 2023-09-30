package main

//go:generate go run .

import (
	"fmt"
	"lab/exp1/src/controllers"
	"lab/exp1/src/db"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title																	Fiber Example API
// @version																v1
// @description														This is a sample swagger for Fiber
// @contact.name													Edy KOFFI
// @contact.email													koffiedy@gmail.com
// @contact.url														https://linkedin.com/in/edy-koffi
// @license.name													Apache 2.0
// @license.url														http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey.name 			ApiKeyAuth
// @securityDefinitions.apikey.in					header
func main() {

	app := fiber.New()

	app.Static("/public", "public")

	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(favicon.New(favicon.Config{File: "./public/favicon.ico"}))
	app.Use(logger.New())

	// Database migration
	db.Db().AutoMigrate(&db.Books{}, &db.User{})

	// Controller initialisation
	controllers.New(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Fatal(app.Listen(port))
}
