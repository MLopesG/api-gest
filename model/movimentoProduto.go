package model

import "github.com/go-playground/validator/v10"

type MovimentoProduto struct {
	Id             int64             `json:"id"`
	DestinoProduto string            `json:"destino_produto" validate:"required"`
	Quantidade     int64             `json:"quantidade" validate:"required"`
	ProdutoId      int64             `json:"produto_id" validate:"required"`
	UsuarioId      int64             `json:"usuario_id" validate:"required"`
	VeiculoId      int64             `json:"veiculo_id"`
	IsEntrada      bool              `json:"is_entrada"`
	CreatedAt      DateFormattedTime `json:"created_at"`
	UpdatedAt      DateFormattedTime `json:"updated_at"`
}

type MovimentoEstoque struct {
	Id                  int64             `json:"id"`
	DestinoProduto      string            `json:"destino_produto"`
	Produto             string            `json:"produto"`
	Usuario             string            `json:"usuario"`
	VeiculoPlaca        string            `json:"veiculo_placa"`
	TipoMovimento       string            `json:"tipo_movimento"`
	Quantidade          int64             `json:"quantidade"`
	QuantidadeEmEstoque int64             `json:"quantidade_em_estoque"`
	RegistradoEm        DateFormattedTime `json:"registrado_em"`
}

func ValidateMovimentoProduto(movimentoProduto MovimentoProduto) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(movimentoProduto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Input = err.StructField()
			element.Value = "Campo é obrigatório!"
			errors = append(errors, &element)
		}
	}
	return errors
}
