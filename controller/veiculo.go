package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func Veiculos(c *fiber.Ctx) error {
	db := database.DB
	var veiculos []model.VeiculoCategoria

	db.Raw(`
		select veiculo.*, 
			categoria.nome as nome_categoria, 
			categoria.is_categoria_produto, 
			categoria.is_categoria_manutencao, 
			categoria.is_categoria_veiculo 
		from veiculo
		left join categoria on categoria.id = veiculo.categoria_id
		order by categoria.id desc
	`).Scan(&veiculos)

	if len(veiculos) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhum veiculo cadastrado", "veiculo": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "veiculos": veiculos})
}

func CadastrarVeiculo(c *fiber.Ctx) error {
	db := database.DB
	veiculo := new(model.Veiculo)

	err := c.BodyParser(veiculo)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}

	errors := model.ValidateVeiculo(*veiculo)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	err = db.Table("veiculo").Create(&veiculo).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel cadastrar veiculo", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Veiculo cadastrado com sucesso!", "veiculo": veiculo})
}

func Veiculo(c *fiber.Ctx) error {
	db := database.DB
	var veiculo model.Veiculo

	id := c.Params("id")

	db.Table("veiculo").Find(&veiculo, "id = ?", id)

	if veiculo.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do veiculo não informado!", "veiculo": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Veiculo encontrado!", "veiculo": veiculo})
}

func DeletarVeiculo(c *fiber.Ctx) error {
	db := database.DB
	var veiculo model.Veiculo

	id := c.Params("id")

	db.Table("veiculo").Find(&veiculo, "id = ?", id)

	if veiculo.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do veiculo não informado!", "veiculo": nil})
	}

	err := db.Table("veiculo").Delete(&veiculo, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar cadastro do usuário", "veiculo": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Veiculo deletado com sucesso!", "veiculo": veiculo})
}

func AlterarVeiculo(c *fiber.Ctx) error {
	db := database.DB
	veiculo  := new(model.Veiculo)

	id := c.Params("id")

	db.Table("veiculo").Find(&veiculo, "id = ?", id)

	if veiculo.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do veiculo não informado!", "veiculo": nil})
	}

	var veiculoAlterar model.VeiculoUpdate

	err := c.BodyParser(&veiculoAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel realizar cadastro, verifique sua senha.", "error": err})
	}

	veiculo.Placa = veiculoAlterar.Placa
	veiculo.Descricao = veiculoAlterar.Descricao
	veiculo.CategoriaId = veiculoAlterar.CategoriaId
	veiculo.IsServico = veiculoAlterar.IsServico

	errors := model.ValidateVeiculo(*veiculo)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	db.Table("veiculo").Save(&veiculo)

	return c.JSON(fiber.Map{"status": true, "message": "Cadastro do veiculo foi alterado com sucesso!"})
}
