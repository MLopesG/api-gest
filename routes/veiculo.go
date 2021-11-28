package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasVeiculo(router fiber.Router) {

	routerFiber := router.Group("/veiculos")

	routerFiber.Post("/", controller.CadastrarVeiculo)
	routerFiber.Get("/", controller.Veiculos)
	routerFiber.Get("/:id", controller.Veiculo)
	routerFiber.Put("/:id", controller.AlterarVeiculo)
	routerFiber.Delete("/:id", controller.DeletarVeiculo)
}
