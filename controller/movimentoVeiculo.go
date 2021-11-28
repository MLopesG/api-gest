package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func MovimentosVeiculos(c *fiber.Ctx) error {
	db := database.DB
	var movimentos []model.MovimentoVeiculo

	db.Raw(`
		select
			movimento_veiculo.id, 
			veiculo.placa, 
			usuario.nome as usuario,
			movimento_veiculo.data_saida_entrada,
			movimento_veiculo.data_retorno_chegada,
			movimento_veiculo.tipo_movimento,
			movimento_veiculo.km_saida,
			movimento_veiculo.km_entrada,
			movimento_veiculo.created_at as registrado_em
		from movimento_veiculo
		left join usuario on usuario.id = movimento_veiculo.usuario_id
		left join veiculo on veiculo.id = movimento_veiculo.veiculo_id
		order by movimento_veiculo.id desc
	`).Scan(&movimentos)

	if len(movimentos) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhuma movimentação registrada!", "movimentos": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "movimentos": movimentos})
}

func RegistrarMovimentoVeicular(c *fiber.Ctx) error {
	db := database.DB
	movimento := new(model.MovimentoVeiculo)

	err := c.BodyParser(movimento)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}

	errors := model.ValidateMovimentoVeiculo(*movimento)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	err = db.Table("movimento_veiculo").Create(&movimento).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel registrar movimentação veicular!", "error": err})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Movimenteção veicular registrada com sucesso!", "movimento": movimento})
}

func Movimentacao(c *fiber.Ctx) error {
	db := database.DB
	var movimento model.MovimentoVeiculo

	id := c.Params("id")

	db.Table("movimento_veiculo").Find(&movimento, "id = ?", id)

	if movimento.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da movimentação não informada!", "manutencao": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Movimentação encontrada!", "movimento": movimento})
}

func DeletarMovimentacaoVeicular(c *fiber.Ctx) error {
	db := database.DB
	var movimento model.MovimentoVeiculo

	id := c.Params("id")

	db.Table("movimento_veiculo").Find(&movimento, "id = ?", id)

	if movimento.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da movimentação não informada!", "movimento": nil})
	}

	err := db.Table("movimento_veiculo").Delete(&movimento, "id = ?", id).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Falha ao deletar registro de movimentação!", "movimento": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Registro de movimentação deletada com sucesso!", "movimento": movimento})
}

func AlterarRegistroMovimentacaoVeicular(c *fiber.Ctx) error {
	db := database.DB
	movimento := new(model.MovimentoVeiculo)

	id := c.Params("id")

	db.Table("movimento_veiculo").Find(&movimento, "id = ?", id)

	if movimento.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Identificador da movimentação não informada!", "movimento": nil})
	}

	var movimentoAlterar model.MovimentoVeiculoUpdate

	err := c.BodyParser(&movimentoAlterar)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel alterar registro!", "error": err})
	}

	movimento.DataSaidaEntrada = movimentoAlterar.DataSaidaEntrada
	movimento.DataRetornoChegada = movimentoAlterar.DataRetornoChegada
	movimento.TipoMovimento = movimentoAlterar.TipoMovimento
	movimento.KmSaida = movimentoAlterar.KmSaida
	movimento.KmEntrada = movimentoAlterar.KmEntrada
	movimento.VeiculoId = movimentoAlterar.VeiculoId
	movimento.UsuarioId = movimentoAlterar.UsuarioId
	
	errors := model.ValidateMovimentoVeiculo(*movimento)

    if errors != nil {
	   return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
    }

	db.Table("movimento_veiculo").Save(&movimento)

	return c.JSON(fiber.Map{"status": true, "message": "Registro de movimentação veicular foi alterado com sucesso!"})
}
