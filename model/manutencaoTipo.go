package model

import "github.com/go-playground/validator/v10"

type ManutencaoTipo struct {
	Id                int64             `json:"id"`
	Descricao         string            `json:"descricao" validate:"required"`
	KmPrevisto        int64             `json:"km_previsto" validate:"required"`
	IntervaloPrevisto int64             `json:"intervalo_previsto" validate:"required"`
	CategoriaId       int64             `json:"categoria_id" validate:"required"`
	CreatedAt         DateFormattedTime `json:"created_at"`
	UpdatedAt         DateFormattedTime `json:"updated_at"`
}

type ManutencaoTipoUpdate struct {
	Descricao         string `json:"descricao"`
	KmPrevisto        int64  `json:"km_previsto"`
	IntervaloPrevisto int64  `json:"intervalo_previsto"`
	CategoriaId       int64  `json:"categoria_id"`
}

type ManutencaoTipoCategoria struct {
	Id                int64  `json:"id"`
	Descricao         string `json:"descricao"`
	KmPrevisto        int64  `json:"km_previsto"`
	IntervaloPrevisto int64  `json:"intervalo_previsto"`
	CategoriaId       int64  `json:"categoria_id"`
	Categoria         string `json:"categoria"`
}

func ValidateManutencaoTipo(tipo ManutencaoTipo) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(tipo)
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
