package model

import "github.com/go-playground/validator/v10"

type Categoria struct {
	Id                    int               `json:"id"`
	Nome                  string            `json:"nome"  validate:"required"`
	IsCategoriaProduto    bool              `json:"is_categoria_produto"`
	IsCategoriaManutencao bool              `json:"is_categoria_manutencao"`
	IsCategoriaVeiculo    bool              `json:"is_categoria_veiculo"`
	CreatedAt             DateFormattedTime `json:"created_at"`
	UpdatedAt             DateFormattedTime `json:"updated_at"`
}

type CategoriaUpdate struct {
	Nome                  string `json:"nome"`
	IsCategoriaProduto    bool   `json:"is_categoria_produto"`
	IsCategoriaManutencao bool   `json:"is_categoria_manutencao"`
	IsCategoriaVeiculo    bool   `json:"is_categoria_veiculo"`
}

func ValidateCategoria(categoria Categoria) []*ErrorResponse {
    var errors []*ErrorResponse
    validate := validator.New()
    err := validate.Struct(categoria)
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