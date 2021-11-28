package routes

import (
	"gestfro/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasUsuario(router fiber.Router) {

	routerFiber := router.Group("/usuarios")

	routerFiber.Post("/", controller.Cadastrar)
	routerFiber.Get("/", controller.Usuarios)
	routerFiber.Get("/:id", controller.Usuario)
	routerFiber.Put("/:id", controller.Alterar)
	routerFiber.Delete("/:id", controller.Deletar)
}
