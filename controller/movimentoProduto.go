package controller

import (
	"gestfro/database"
	"gestfro/model"

	"github.com/gofiber/fiber/v2"
)

func MovimentosEstoque(c *fiber.Ctx) error {
	db := database.DB
	var movimentos []model.MovimentoEstoque

	db.Raw(`
		select 
			movimento_produto.id,
			movimento_produto.destino_produto,
			produto.nome as produto,
			usuario.nome as usuario,
			veiculo.placa as veiculo_placa,
			(
			case 
				when movimento_produto.is_entrada then 'Entrada de Produto.'
				else 'Saida de Produto.'
			end
			) as tipo_movimento,
			(
			case 
				when produto.quantidade between 1 and 3 then 'Baixo.'
				when produto.quantidade = 0 then 'Em falta.'
				else 'Normal.'
			end
			) as status_produto,
			movimento_produto.quantidade,
			produto.quantidade as quantidade_em_estoque,
			movimento_produto.created_at as registrado_em
		from  movimento_produto
		inner join usuario on usuario.id = movimento_produto.usuario_id
		inner join produto on produto.id = movimento_produto.produto_id
		left join veiculo on veiculo.id = movimento_produto.veiculo_id
		order by movimento_produto.created_at desc;
	`).Scan(&movimentos)

	if len(movimentos) == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Nenhuma movimentação registrada!", "movimentos": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "movimentos": movimentos})
}

func RegistrarSaidaEntradaProduto(c *fiber.Ctx) error {
	db := database.DB

	var produto model.Produto

	movimento := new(model.MovimentoProduto)

	err := c.BodyParser(movimento)

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Corpo inválido!", "error": err})
	}

	errors := model.ValidateMovimentoProduto(*movimento)

	if errors != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Preenche os campos corretamente", "errors": errors})
	}

	/// Consultar produto
	db.Table("produto").Find(&produto, "id = ?", movimento.ProdutoId)

	if produto.Id == 0 {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Produto não foi encontrado!", "produto": nil})
	}

	if !movimento.IsEntrada {
		if produto.Quantidade < movimento.Quantidade || movimento.Quantidade < 0 {
			return c.Status(417).JSON(fiber.Map{"status": false, "message": "Quantidade selecionada não está disponivel no estoque!", "produto": produto})
		} else {
			/// Descontar produto retirado do estoque
			produto.Quantidade -= movimento.Quantidade
		}
	} else {
		/// Somar quantidade de entrada no produto
		produto.Quantidade += movimento.Quantidade
	}

	err = db.Table("movimento_produto").Create(&movimento).Error

	if err != nil {
		return c.Status(417).JSON(fiber.Map{"status": false, "message": "Não foi possivel registrar movimento do produto!", "error": err})
	}

	///Atualizar cadastro do produto com nova quantidade
	db.Table("produto").Save(&produto)

	return c.JSON(fiber.Map{"status": true, "message": "Movimento registrado com sucesso!", "movimento": movimento})
}
