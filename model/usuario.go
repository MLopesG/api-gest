package model

import "github.com/go-playground/validator/v10"

type Usuario struct {
	Id        int64             `json:"id"`
	Nome      string            `json:"nome" validate:"required"`
	IsAtivo   bool              `json:"is_ativo" validate:"required"`
	Senha     string            `json:"senha" validate:"required"`
	Cpf       string            `json:"cpf" validate:"required"`
	Email     string            `json:"email" validate:"required"`
	CreatedAt DateFormattedTime `json:"created_at"`
	UpdatedAt DateFormattedTime `json:"updated_at"`
}

type UsuarioUpdate struct {
	Nome    string `json:"nome"`
	IsAtivo bool   `json:"is_ativo"`
	Senha   string `json:"senha"`
	Cpf     string `json:"cpf"`
	Email   string `json:"email"`
}

func ValidateUsuario(usuario Usuario) []*ErrorResponse {
    var errors []*ErrorResponse
    validate := validator.New()
    err := validate.Struct(usuario)
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