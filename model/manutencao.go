package model

import "github.com/go-playground/validator/v10"

type Manutencao struct {
	Id        int64    `json:"id"`
	KmAtual   int64   `json:"km_atual" validate:"required"`
	Descricao string     `json:"descricao" validate:"required"`
	VeiculoId int64   `json:"veiculo_id" validate:"required"`
	CategoriaId  int64   `json:"categoria_id" validate:"required"`
	UsuarioId  int64   `json:"usuario_id" validate:"required"`
	CreatedAt DateFormattedTime `json:"created_at"`
	UpdatedAt DateFormattedTime `json:"updated_at"`
}

type ManutencaoUpdate struct {
	KmAtual  	int64   `json:"km_atual"`
	Descricao   string `json:"descricao"`
	VeiculoId   int64  `json:"veiculo_id"`
	CategoriaId int64  `json:"categoria_id"`
	UsuarioId   int64  `json:"usuario_id"`
}

type ManutencaoVeiculoUsuarioCategoria struct {
	Id     int     `json:"id"`
	Placa  string  `json:"placa"`
	Usuario   string  `json:"usuario"`
	KmAtual  	int64   `json:"km_atual"`
	Descricao   string `json:"descricao"`
	NomeCategoria  string `json:"nome_categoria"`
	RegistradoEm DateFormattedTime `json:"registrado_em"`
}

func ValidateManutencao(manutencao Manutencao) []*ErrorResponse {
    var errors []*ErrorResponse
    validate := validator.New()
    err := validate.Struct(manutencao)
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