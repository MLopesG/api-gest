package main

import (
	"gestfro/database"
	"gestfro/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/monitor", monitor.New())
	
	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
