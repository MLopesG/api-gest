package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasProduto(router fiber.Router) {

	routerFiber := router.Group("/produtos")

	routerFiber.Post("/", controller.CadastrarProduto)
	routerFiber.Get("/", controller.Produtos)
	routerFiber.Get("/:id", controller.Produto)
	routerFiber.Put("/:id", controller.AlterarProduto)
	routerFiber.Delete("/:id", controller.DeletarProduto)
}
