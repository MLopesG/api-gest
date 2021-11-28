package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func Categorias(c *fiber.Ctx) error {
	db := database.DB
	var categorias []model.Categoria

	db.Table("categoria").Order("nome desc").Find(&categorias)

	if len(categorias) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhum categoria cadastrada", "categorias": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "categorias": categorias})
}

func CadastrarCategoria(c *fiber.Ctx) error {
	db := database.DB
	categoria := new(model.Categoria)

	err := c.BodyParser(categoria)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}
	
	errors := model.ValidateCategoria(*categoria)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	err = db.Table("categoria").Create(&categoria).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel cadastrar nova categoria", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Categoria cadastrada com sucesso!", "categoria": categoria})
}

func Categoria(c *fiber.Ctx) error {
	db := database.DB
	var categoria model.Categoria

	id := c.Params("id")

	db.Table("categoria").Find(&categoria, "id = ?", id)

	if categoria.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da categoria não informado!", "categoria": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Categoria encontrado!", "categoria": categoria})
}

func DeletarCategoria(c *fiber.Ctx) error {
	db := database.DB
	var categoria model.Categoria

	id := c.Params("id")

	db.Table("categoria").Find(&categoria, "id = ?", id)

	if categoria.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da categoria não informado!", "categoria": nil})
	}

	err := db.Table("categoria").Delete(&categoria, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar cadastro da categoria!", "usuario": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Categoria deletada com sucesso!", "categoria": categoria})
}

func AlterarCategoria(c *fiber.Ctx) error {
	db := database.DB
	categoria := new(model.Categoria)

	id := c.Params("id")

	db.Table("categoria").Find(&categoria, "id = ?", id)

	if categoria.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da categoria não informado!", "categoria": nil})
	}

	var categoriaAlterar model.CategoriaUpdate

	err := c.BodyParser(&categoriaAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "usuario": err})
	}

	categoria.Nome = categoriaAlterar.Nome
	categoria.IsCategoriaProduto = categoriaAlterar.IsCategoriaProduto
	categoria.IsCategoriaManutencao = categoriaAlterar.IsCategoriaManutencao
	categoria.IsCategoriaVeiculo = categoriaAlterar.IsCategoriaVeiculo

	errors := model.ValidateCategoria(*categoria)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	db.Table("categoria").Save(&categoria)

	return c.JSON(fiber.Map{"status": true, "message": "Cadastro de categoria foi alterado com sucesso!"})
}
