package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasMovimentoVeiculo(router fiber.Router) {

	routerFiber := router.Group("/movimentos-veiculares")

	routerFiber.Post("/", controller.RegistrarMovimentoVeicular)
	routerFiber.Get("/", controller.MovimentosVeiculos)
	routerFiber.Get("/:id", controller.Movimentacao)
	routerFiber.Put("/:id", controller.AlterarRegistroMovimentacaoVeicular)
	routerFiber.Delete("/:id", controller.DeletarMovimentacaoVeicular)
}
