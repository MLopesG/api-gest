package routes

import (
	"gestfro/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasCategoria(router fiber.Router) {

	routerFiber := router.Group("/categorias")

	routerFiber.Post("/", handler.CadastrarCategoria)
	routerFiber.Get("/", handler.Categorias)
	routerFiber.Get("/:id", handler.Categoria)
	routerFiber.Put("/:id", handler.AlterarCategoria)
	routerFiber.Delete("/:id", handler.DeletarCategoria)
}
