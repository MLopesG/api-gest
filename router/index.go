package router

import (
	"gestfro/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1/gestfro", logger.New())

	routes.SetupRotasUsuario(api)
	routes.SetupRotasCategoria(api)
}
