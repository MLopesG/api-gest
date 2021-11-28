package controller

import (
	"gestfro/database"
	"gestfro/model"
	"github.com/gofiber/fiber/v2"
)

func Usuarios(c *fiber.Ctx) error {
	db := database.DB
	var usuarios []model.Usuario

	db.Table("usuario").Order("nome desc").Find(&usuarios)

	if len(usuarios) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhum usuário cadastrado", "usuario": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "usuarios": usuarios})
}

func Cadastrar(c *fiber.Ctx) error {
	db := database.DB
	usuario := new(model.Usuario)

	err := c.BodyParser(usuario)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}

	errors := model.ValidateUsuario(*usuario)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	hash, err := HashPassword(usuario.Senha)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel realizar cadastro, verifique sua senha.", "error": err})
	}

	usuario.Senha = hash

	err = db.Table("usuario").Create(&usuario).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel cadastrar usuário", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Usuário cadastrado com sucesso!", "usuario": usuario})
}

func Usuario(c *fiber.Ctx) error {
	db := database.DB
	var usuario model.Usuario

	id := c.Params("id")

	db.Table("usuario").Find(&usuario, "id = ?", id)

	if usuario.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do usuário não informado!", "usuario": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Usuário encontrado!", "usuario": usuario})
}

func Deletar(c *fiber.Ctx) error {
	db := database.DB
	var usuario model.Usuario

	id := c.Params("id")

	db.Table("usuario").Find(&usuario, "id = ?", id)

	if usuario.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do usuário não informado!", "usuario": nil})
	}

	err := db.Table("usuario").Delete(&usuario, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar cadastro do usuário", "usuario": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Usuário deletado com sucesso!", "usuario": usuario})
}

func Alterar(c *fiber.Ctx) error {
	db := database.DB
	usuario := new(model.Usuario)

	id := c.Params("id")

	db.Table("usuario").Find(&usuario, "id = ?", id)

	if usuario.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do usuário não informado!", "usuario": nil})
	}

	var usuarioAlterar model.UsuarioUpdate

	err := c.BodyParser(&usuarioAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "usuario": err})
	}

	hash, err := HashPassword(usuarioAlterar.Senha)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel realizar cadastro, verifique sua senha.", "error": err})
	}

	usuario.Nome = usuarioAlterar.Nome
	usuario.IsAtivo = usuarioAlterar.IsAtivo
	usuario.Senha = hash
	usuario.Cpf = usuarioAlterar.Cpf
	usuario.Email = usuarioAlterar.Email
	
	errors := model.ValidateUsuario(*usuario)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	db.Table("usuario").Save(&usuario)

	return c.JSON(fiber.Map{"status": true, "message": "Cadastro do usuário foi alterado com sucesso!"})
}
