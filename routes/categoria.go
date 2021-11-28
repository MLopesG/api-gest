package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasCategoria(router fiber.Router) {

	routerFiber := router.Group("/categorias")

	routerFiber.Post("/", controller.CadastrarCategoria)
	routerFiber.Get("/", controller.Categorias)
	routerFiber.Get("/:id", controller.Categoria)
	routerFiber.Put("/:id", controller.AlterarCategoria)
	routerFiber.Delete("/:id", controller.DeletarCategoria)
}
