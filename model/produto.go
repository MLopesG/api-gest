package model

import "github.com/go-playground/validator/v10"

type Produto struct {
	Id                 int64             `json:"id"`
	Quantidade         int64             `json:"quantidade"`
	Nome               string            `json:"nome"  validate:"required"`
	DataEntradaEstoque DateFormattedTime `json:"data_entrada_estoque" validate:"required"`
	CategoriaId        int64             `json:"categoria_id" validate:"required"`
	CreatedAt          DateFormattedTime `json:"created_at"`
	UpdatedAt          DateFormattedTime `json:"updated_at"`
}

type ProdutoUpdate struct {
	Nome               string            `json:"nome"`
	Quantidade         int64             `json:"id"`
	DataEntradaEstoque DateFormattedTime `json:"data_entrada_estoque"`
	CategoriaId        int64             `json:"categoria_id"`
}

type ProdutoCategoria struct {
	Id            int64  `json:"id"`
	Quantidade    int64  `json:"quantidade"`
	Nome          string `json:"nome"`
	CategoriaId   int64  `json:"categoria_id"`
	NomeCategoria string `json:"nome_categoira"`
	StatusProduto string `json:"status_produto"`
}

func ValidateProduto(produto Produto) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(produto)
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
