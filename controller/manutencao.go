package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func Manutencoes(c *fiber.Ctx) error {
	db := database.DB
	var manutecoes []model.Manuntencoes

	db.Raw(`
		select
			m.id,
			m.km_atual,
			m.descricao, 
			m.valor_pago,
			mt.descricao as categoria_manutencao,
			(
				case
					when m.is_finalizado then 'Concluido.'
					when m.is_andamento then 'Em Andamento.'
					when m.is_cancelado then 'Cancelado.'
					else 'Não identificado.'
				end
			) as status_manutencao,
			v.id as id_veiculo_manutencao,
			v.placa as veiculo_manutencao,
			(
				case
					when vs.descricao is null then 'Veiculo não substituido.'
					else vs.placa
				end 
			) as veiculo_substituto,
			vs.id as id_veiculo_substituto,
			u.id as usuario_id,
			u.nome as usuario,
			m.cancelado_em,
			m.created_at as registrado_em
		from manutencao m
		inner join veiculo v on v.id = m.veiculo_id 
		left  join usuario u on u.id = m.usuario_id 
		inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
		left join veiculo vs on vs.id = m.veiculo_id_temporario
		order by m.id desc;
	`).Scan(&manutecoes)

	if len(manutecoes) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhuma Manutenção Registrada!", "manutecoes": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "manutecoes": manutecoes})
}

func CadastrarNovaManutencao(c *fiber.Ctx) error {
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
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel registrar manutenção!", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Manutenção registrada com sucesso!", "manutencao": manutencao})
}

func Manutencao(c *fiber.Ctx) error {
	db := database.DB
	var manutencao model.Manutencao

	id := c.Params("id")

	db.Table("manutencao").Find(&manutencao, "id = ?", id)

	if manutencao.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da manutenção não informada!", "tipo": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Manutencão encontrado!", "manutencao": manutencao})
}

func DeletarManutencao(c *fiber.Ctx) error {
	db := database.DB
	var manutencao model.Manutencao

	id := c.Params("id")

	db.Table("manutencao").Find(&manutencao, "id = ?", id)

	if manutencao.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da manutenção não informada!", "tipo": nil})
	}

	err := db.Table("manutencao").Delete(&manutencao, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar registro!", "movimento": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Manutenção excluida com sucesso!", "manutencao": manutencao})
}

func AlterarManutencao(c *fiber.Ctx) error {
	db := database.DB
	manutencao := new(model.Manutencao)

	id := c.Params("id")

	db.Table("manutencao").Find(&manutencao, "id = ?", id)

	if manutencao.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da manutenção não informada!", "tipo": nil})
	}
	var manutencaoAlterar model.ManutencaoUpdate

	err := c.BodyParser(&manutencaoAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel alterar registro!", "error": err})
	}

	manutencao.KmAtual = manutencaoAlterar.KmAtual
	manutencao.Descricao = manutencaoAlterar.Descricao
	manutencao.ValorPago = manutencaoAlterar.ValorPago
	manutencao.VeiculoId = manutencaoAlterar.VeiculoId
	manutencao.UsuarioId = manutencaoAlterar.UsuarioId
	manutencao.UsuarioId = manutencaoAlterar.ManutencaoTipoId
	manutencao.ManutencaoTipoId = manutencaoAlterar.ManutencaoTipoId
	manutencao.IsFinalizado = manutencaoAlterar.IsFinalizado
	manutencao.IsAndamento = manutencaoAlterar.IsAndamento
	manutencao.IsCancelado = manutencaoAlterar.IsCancelado
	manutencao.VeiculoIdTemporario = manutencaoAlterar.VeiculoIdTemporario
	manutencao.CanceladoEm = manutencaoAlterar.CanceladoEm

	errors := model.ValidateManutencao(*manutencao)

	if errors != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
	}

	db.Table("manutencao").Save(&manutencao)

	return c.JSON(fiber.Map{"status": true, "message": "Manutenção foi alterado com sucesso!"})
}
