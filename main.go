package main

import (
	"gestfro/database"
	"gestfro/router"
	"gestfro/config"
	"fmt"

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

	app.Get("/", monitor.New())
	
	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(fmt.Sprintf(":%s", config.Config("PORT")))
}
