package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasManutencao(router fiber.Router) {

	routerFiber := router.Group("/manutencoes")

	routerFiber.Post("/", controller.RegistrarManutencao)
	routerFiber.Get("/", controller.Manutencoes)
	routerFiber.Get("/:id", controller.Manutencao)
	routerFiber.Put("/:id", controller.AlterarRegistroManuntencao)
	routerFiber.Delete("/:id", controller.DeletarManutencao)
}
