package routes

import (
	"gestfro/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasVeiculo(router fiber.Router) {

	routerFiber := router.Group("/veiculos")

	routerFiber.Post("/", handler.CadastrarVeiculo)
	routerFiber.Get("/", handler.Veiculos)
	routerFiber.Get("/:id", handler.Veiculo)
	routerFiber.Put("/:id", handler.AlterarVeiculo)
	routerFiber.Delete("/:id", handler.DeletarVeiculo)
}
