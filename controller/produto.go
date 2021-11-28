package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func Produtos(c *fiber.Ctx) error {
	db := database.DB
	var produtos []model.ProdutoCategoria

	db.Raw(
		`select 
		 	produto.id, 
		 	produto.nome, 
			produto.quantidade, 
			categoria.id as categoria_id,
			categoria.nome as nome_categoria
		 from produto
		 left join categoria on categoria.id = produto.categoria_id
		`).Scan(&produtos)

	if len(produtos) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhum produto cadastrado!", "produtos": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "produtos": produtos})
}

func CadastrarProduto(c *fiber.Ctx) error {
	db := database.DB
	produto := new(model.Produto)

	err := c.BodyParser(produto)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}
	
	errors := model.ValidateProduto(*produto)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	err = db.Table("produto").Create(&produto).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel cadastrar novo produto!", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Produto cadastrado com sucesso!", "produto": produto})
}

func Produto(c *fiber.Ctx) error {
	db := database.DB
	var produto model.Produto

	id := c.Params("id")

	db.Table("produto").Find(&produto, "id = ?", id)

	if produto.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do produto não informado!", "produto": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Produto encontrado!", "produto": produto})
}

func DeletarProduto(c *fiber.Ctx) error {
	db := database.DB
	var produto model.Produto

	id := c.Params("id")

	db.Table("produto").Find(&produto, "id = ?", id)

	if produto.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do produto não informado!", "produto": nil})
	}

	err := db.Table("produto").Delete(&produto, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar cadastro do produto!", "usuario": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Produto deletado com sucesso!", "produto": produto})
}

func AlterarProduto(c *fiber.Ctx) error {
	db := database.DB
	produto := new(model.Produto)

	id := c.Params("id")

	db.Table("produto").Find(&produto, "id = ?", id)

	if produto.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do produto não informado!", "produto": nil})
	}
	
	var produtoAlterar model.ProdutoUpdate

	err := c.BodyParser(&produtoAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "usuario": err})
	}

	produto.Nome = produtoAlterar.Nome
	produto.DataEntradaEstoque = produtoAlterar.DataEntradaEstoque
	produto.CategoriaId = produtoAlterar.CategoriaId

	errors := model.ValidateProduto(*produto)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	db.Table("produto").Save(&produto)

	return c.JSON(fiber.Map{"status": true, "message": "Cadastro de produto foi alterado com sucesso!"})
}
