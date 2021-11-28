package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func Manutencoes(c *fiber.Ctx) error {
	db := database.DB
	var manutencoes []model.ManutencaoVeiculoUsuarioCategoria

	db.Raw(`
		select
			manutencao.id, 
			veiculo.placa, 
			usuario.nome as usuario, 
			manutencao.km_atual, 
			manutencao.descricao, 
			categoria.nome as nome_categoria,
			manutencao.created_at as registrado_em
		from manutencao
		left join categoria on categoria.id = manutencao.categoria_id
		left join usuario on usuario.id = manutencao.usuario_id
		left join veiculo on veiculo.id = manutencao.veiculo_id
		order by manutencao.id desc
	`).Scan(&manutencoes)

	if len(manutencoes) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhum manutenção registrada!", "manutencoes": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "manutencoes": manutencoes})
}

func RegistrarManutencao(c *fiber.Ctx) error {
	db := database.DB
	manutencao := new(model.Manutencao)

	err := c.BodyParser(manutencao)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}

	errors := model.ValidateManutencao(*manutencao)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	err = db.Table("manutencao").Create(&manutencao).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel registrar manuntenção!", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Manutenção registrada com sucesso!", "manutencao": manutencao})
}

func Manutencao(c *fiber.Ctx) error {
	db := database.DB
	var manutencao model.Manutencao

	id := c.Params("id")

	db.Table("manutencao").Find(&manutencao, "id = ?", id)

	if manutencao.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da manutenção não informada!", "manutencao": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Manutenção encontrada!", "manutencao": manutencao})
}

func DeletarManutencao(c *fiber.Ctx) error {
	db := database.DB
	var manutencao model.Manutencao

	id := c.Params("id")

	db.Table("manutencao").Find(&manutencao, "id = ?", id)

	if manutencao.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da manutenção não informada!", "manutencao": nil})
	}

	err := db.Table("manutencao").Delete(&manutencao, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar registro de manutenção!", "manutencao": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Registro de manutenção deletado com sucesso!", "manutencao": manutencao})
}

func AlterarRegistroManuntencao(c *fiber.Ctx) error {
	db := database.DB
	manutencao := new(model.Manutencao)

	id := c.Params("id")

	db.Table("manutencao").Find(&manutencao, "id = ?", id)

	if manutencao.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da manutenção não informada!", "manutencao": nil})
	}

	var manutencaoAlterar model.ManutencaoUpdate

	err := c.BodyParser(&manutencaoAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel alterar registro!", "error": err})
	}

	manutencao.KmAtual = manutencaoAlterar.KmAtual
	manutencao.Descricao = manutencaoAlterar.Descricao
	manutencao.VeiculoId = manutencaoAlterar.VeiculoId
	manutencao.CategoriaId = manutencaoAlterar.CategoriaId
	manutencao.UsuarioId = manutencaoAlterar.UsuarioId
	
	errors := model.ValidateManutencao(*manutencao)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	db.Table("manutencao").Save(&manutencao)

	return c.JSON(fiber.Map{"status": true, "message": "Registro de manutenção alterado com sucesso!"})
}
