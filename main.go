package main

//go:generate go run .

import (
	"lab/exp1/src/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title			Fiber Example API
// @version		v1
// @description	This is a sample swagger for Fiber
// @contact.name	Edy KOFFI
// @contact.email	koffiedy@gmail.com
// @contact.url	koffiedy@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	app := fiber.New()

	app.Static("/public", "./public")

	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{File: "./public/favicon.ico"}))
	app.Use(logger.New())

	controllers.New(app)

	log.Fatal(app.Listen(":8080"))
}
