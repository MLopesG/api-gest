package model

import "github.com/go-playground/validator/v10"

type Veiculo struct {
	Id             int64             `json:"id"`
	Placa          string            `json:"placa"  validate:"required"`
	Descricao      string            `json:"descricao"  validate:"required"`
	CategoriaId    int64             `json:"categoria_id"  validate:"required"`
	IsServico      bool              `json:"is_servico"`
	IsReserva      bool              `json:"is_reserva"`
	IsDisponivel   bool              `json:"is_disponivel"`
	IsIndisponivel bool              `json:"is_indisponivel"`
	IsSubstituido  bool              `json:"is_substituido"`
	CreatedAt      DateFormattedTime `json:"created_at"`
	UpdatedAt      DateFormattedTime `json:"updated_at"`
}

type VeiculoCategoria struct {
	Id             int64  `json:"id"`
	Placa          string `json:"placa"`
	CategoriaId    int64  `json:"categoria_id"`
	NomeCategoria  string `json:"nome_categoria"`
	Descricao      string `json:"descricao"`
	StatusOperacao string `json:"status_operacao"`
	TipoVeiculo    string `json:"tipo_veiculo"`
}

type VeiculoUpdate struct {
	Placa          string `json:"placa"`
	Descricao      string `json:"descricao"`
	CategoriaId    int64  `json:"categoria_id"`
	IsServico      bool   `json:"is_servico"`
	IsReserva      bool   `json:"is_reserva"`
	IsDisponivel   bool   `json:"is_disponivel"`
	IsIndisponivel bool   `json:"is_indisponivel"`
	IsSubstituido  bool   `json:"is_substituido"`
}

func ValidateVeiculo(veiculo Veiculo) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(veiculo)
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
