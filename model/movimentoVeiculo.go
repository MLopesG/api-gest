package model

import "github.com/go-playground/validator/v10"

type MovimentoVeiculo struct {
	Id                 int64             `json:"id"`
	DataSaidaEntrada   DateFormattedTime `json:"data_saida_entrada"`
	DataRetornoChegada DateFormattedTime `json:"data_retorno_chegada"`
	TipoMovimento      string            `json:"tipo_movimento" validate:"required"`
	KmSaida            int64             `json:"km_saida"`
	KmEntrada          int64             `json:"km_entrada"`
	VeiculoId          int64             `json:"veiculo_id" validate:"required"`
	UsuarioId          int64             `json:"usuario_id" validate:"required"`
	CreatedAt          DateFormattedTime `json:"created_at"`
	UpdatedAt          DateFormattedTime `json:"updated_at"`
}

type MovimentoVeiculoUpdate struct {
	DataSaidaEntrada   DateFormattedTime `json:"data_saida_entrada"`
	DataRetornoChegada DateFormattedTime `json:"data_retorno_chegada"`
	TipoMovimento      string            `json:"tipo_movimento"`
	KmSaida            int64             `json:"km_saida"`
	KmEntrada          int64             `json:"km_entrada"`
	VeiculoId          int64             `json:"veiculo_id"`
	UsuarioId          int64             `json:"usuario_id"`
}

type MovimentoVeiculoUsuario struct {
	Id                 int               `json:"id"`
	Placa              string            `json:"placa"`
	Usuario            string            `json:"usuario"`
	DataSaidaEntrada   DateFormattedTime `json:"data_saida_entrada"`
	DataRetornoChegada DateFormattedTime `json:"data_retorno_chegada"`
	TipoMovimento      string            `json:"tipo_movimento"`
	KmSaida            int64             `json:"km_saida"`
	KmEntrada          int64             `json:"km_entrada"`
	RegistradoEm       DateFormattedTime `json:"registrado_em"`
}

func ValidateMovimentoVeiculo(movimento MovimentoVeiculo) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(movimento)
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
