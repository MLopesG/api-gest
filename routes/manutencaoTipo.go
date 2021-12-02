package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasTiposManutencao(router fiber.Router) {

	routerFiber := router.Group("/tipo-manutencoes")

	routerFiber.Post("/", controller.CadastrarNovoTipoManutencao)
	routerFiber.Get("/", controller.Tipos)
	routerFiber.Get("/:id", controller.TipoManutencao)
	routerFiber.Put("/:id", controller.AlterarRegistroTipoManutencao)
	routerFiber.Delete("/:id", controller.DeletarTipoManutencao)
}
