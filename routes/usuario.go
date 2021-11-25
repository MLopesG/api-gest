package routes

import (
	"gestfro/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRotasUsuario(router fiber.Router) {

	routerFiber := router.Group("/usuarios")

	routerFiber.Post("/", handler.Cadastrar)
	routerFiber.Get("/", handler.Usuarios)
	routerFiber.Get("/:id", handler.Usuario)
	routerFiber.Put("/:id", handler.Alterar)
	routerFiber.Delete("/:id", handler.Deletar)
}
