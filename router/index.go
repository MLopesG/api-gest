package router

import (
	"gestfro/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	routes.SetupRotasUsuario(api)
	routes.SetupRotasCategoria(api)
	routes.SetupRotasVeiculo(api)
	routes.SetupRotasProduto(api)
	routes.SetupRotasMovimentoVeiculo(api)
	routes.SetupRotasMovimentoProduto(api)
	routes.SetupRotasTiposManutencao(api)
	routes.SetupRotasManutencao(api)
}
