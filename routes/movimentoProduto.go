package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasMovimentoProduto(router fiber.Router) {

	routerFiber := router.Group("/movimentos-produtos")

	routerFiber.Post("/", controller.RegistrarSaidaEntradaProduto)
	routerFiber.Get("/:veiculo?", controller.MovimentosEstoque)
}
