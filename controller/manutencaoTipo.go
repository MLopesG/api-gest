package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func Tipos(c *fiber.Ctx) error {
	db := database.DB
	var tipos []model.ManutencaoTipoCategoria

	db.Raw(`
		select 
			t.id,
			t.descricao,
			t.km_previsto, 
			t.intervalo_previsto,
			c.id,
			c.nome as categoria
		from manutencao_tipo t
		inner join categoria c on c.id = t.categoria_id 
		order by t.descricao;
	`).Scan(&tipos)

	if len(tipos) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhuma classifciação de manutenção foi cadastrada!", "tipos": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "tipos": tipos})
}

func CadastrarNovoTipoManutencao(c *fiber.Ctx) error {
	db := database.DB
	tipo := new(model.ManutencaoTipo)

	err := c.BodyParser(tipo)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}

	errors := model.ValidateManutencaoTipo(*tipo)

	if errors != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
	}

	err = db.Table("manutencao_tipo").Create(&tipo).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel cadastrar cadastro!", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Cadastro realizado com sucesso!", "tipo": tipo})
}

func TipoManutencao(c *fiber.Ctx) error {
	db := database.DB
	var tipo model.ManutencaoTipo

	id := c.Params("id")

	db.Table("manutencao_tipo").Find(&tipo, "id = ?", id)

	if tipo.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do tipo não informada!", "tipo": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Tipo manutencão encontrado!", "tipo": tipo})
}

func DeletarTipoManutencao(c *fiber.Ctx) error {
	db := database.DB
	var tipo model.ManutencaoTipo

	id := c.Params("id")

	db.Table("manutencao_tipo").Find(&tipo, "id = ?", id)

	if tipo.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do tipo não informada!", "tipo": nil})
	}

	err := db.Table("manutencao_tipo").Delete(&tipo, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar registro!", "movimento": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Tipo removido com sucesso!", "tipo": tipo})
}

func AlterarRegistroTipoManutencao(c *fiber.Ctx) error {
	db := database.DB
	tipo := new(model.ManutencaoTipo)

	id := c.Params("id")

	db.Table("manutencao_tipo").Find(&tipo, "id = ?", id)

	if tipo.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador do tipo não informada!", "tipo": nil})
	}

	var tipoAlterar model.ManutencaoTipoUpdate

	err := c.BodyParser(&tipoAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel alterar registro!", "error": err})
	}

	tipo.Descricao = tipoAlterar.Descricao
	tipo.KmPrevisto = tipoAlterar.KmPrevisto
	tipo.IntervaloPrevisto = tipoAlterar.IntervaloPrevisto
	tipo.CategoriaId = tipoAlterar.CategoriaId

	errors := model.ValidateManutencaoTipo(*tipo)

	if errors != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
	}

	db.Table("manutencao_tipo").Save(&tipo)

	return c.JSON(fiber.Map{"status": true, "message": "Cadastro foi alterado com sucesso!"})
}
