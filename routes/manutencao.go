package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasManutencao(router fiber.Router) {

	routerFiber := router.Group("/manutencoes")

	routerFiber.Post("/", controller.CadastrarNovaManutencao)
	routerFiber.Post("/cancelar/:id", controller.CancelarManutencao)
	routerFiber.Get("/", controller.Manutencoes)
	routerFiber.Get("/:id", controller.Manutencao)
	routerFiber.Put("/:id", controller.AlterarManutencao)
	routerFiber.Delete("/:id", controller.DeletarManutencao)
}
