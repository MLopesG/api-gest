package main

import (
	"gestfro/database"
	"gestfro/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
