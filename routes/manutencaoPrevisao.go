package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasManutencoesPrevisoes(router fiber.Router) {

	routerFiber := router.Group("/manutencoes-previsoes")

	routerFiber.Get("/", controller.Previsoes)
	routerFiber.Get("/dia", controller.PrevisoesManutencaoDia)
	routerFiber.Get("/detalhar/:id", controller.DetalharPrevisaoManutencao)
}
