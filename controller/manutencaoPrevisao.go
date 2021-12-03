package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func DetalharPrevisaoManutencao(c *fiber.Ctx) error {
	db := database.DB
	var previsoes []model.ManutencaoPrevisaoDetalhada

	id := c.Params("id")

	db.Raw(`
			select 
			mp.id,
			mp.km_previsao,
			mp.data_previsao,
			mt.descricao as categoria_manutencao,
			(
				case
					when mp.is_aprovado then 'Aprovado.'
					else 'Cancelado.'
				end
			) as status_previsao,
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
			m.created_at as manutencao_realizado,
			mp.created_at as previsao_aberta_em,
			mp.updated_at as ultima_atualizacao_registro
		FROM manutencao_previsao mp
		inner join manutencao m on m.id = mp.ultima_manutencao_id 
		inner join veiculo v on v.id = m.veiculo_id 
		left  join usuario u on u.id = m.usuario_id 
		inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
		left join veiculo vs on vs.id = m.veiculo_id_temporario
		where mp.id = ?;
	`, id).Scan(&previsoes)

	if len(previsoes) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhuma previsão gerada!", "previsoes": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "previsoes": previsoes})
}

func PrevisoesManutencaoDia(c *fiber.Ctx) error {
	db := database.DB
	var previsoes []model.ManutencaoPrevisaoList

	db.Raw(`
			select 
			mp.id,
			mp.km_previsao,
			mp.data_previsao,
			mt.descricao as categoria_manutencao,
			(
				case
					when mp.is_aprovado then 'Aprovado.'
					else 'Cancelado.'
				end
			) as status_previsao,
			m.descricao, 
			v.placa as veiculo_manutencao,
			mp.created_at as previsao_aberta_em,
			mp.updated_at as ultima_atualizacao_registro
		FROM manutencao_previsao mp
		inner join manutencao m on m.id = mp.ultima_manutencao_id 
		inner join veiculo v on v.id = m.veiculo_id 
		inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
		where date(mp.created_at) = current_date
		and m.is_cancelado = false
		and mp.is_cancelado = false
		order by mp.created_at desc;
	`).Scan(&previsoes)

	if len(previsoes) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhuma previsão gerada!", "previsoes": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "previsoes": previsoes})
}

func Previsoes(c *fiber.Ctx) error {
	db := database.DB
	var previsoes []model.ManutencaoPrevisaoList

	db.Raw(`
		select 
			mp.id,
			mp.km_previsao,
			mp.data_previsao,
			mt.descricao as categoria_manutencao,
			(
				case
					when mp.is_aprovado then 'Aprovado.'
					else 'Cancelado.'
				end
			) as status_previsao,
			m.descricao, 
			v.placa as veiculo_manutencao,
			mp.created_at as previsao_aberta_em,
			mp.updated_at as ultima_atualizacao_registro
		FROM manutencao_previsao mp
		inner join manutencao m on m.id = mp.ultima_manutencao_id 
		inner join veiculo v on v.id = m.veiculo_id 
		inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
		order by mp.created_at desc;
	`).Scan(&previsoes)

	if len(previsoes) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhuma previsão gerada!", "previsoes": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "previsoes": previsoes})
}
