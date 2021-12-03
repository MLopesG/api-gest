package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Manutencao struct {
	Id                  int64             `json:"id"`
	KmAtual             int64             `json:"km_atual" validate:"required"`
	Descricao           string            `json:"descricao" validate:"required"`
	ValorPago           float32           `json:"valor_pago"`
	VeiculoId           int64             `json:"veiculo_id" validate:"required"`
	UsuarioId           int64             `json:"usuario_id" validate:"required"`
	ManutencaoTipoId    int64             `json:"manutencao_tipo_id" validate:"required"`
	IsFinalizado        bool              `json:"is_finalizado"`
	IsAndamento         bool              `json:"is_andamento"`
	IsCancelado         bool              `json:"is_cancelado"`
	VeiculoIdTemporario int64             `json:"veiculo_id_temporario" `
	CanceladoEm         time.Time         `json:"cancelado_em"`
	CreatedAt           DateFormattedTime `json:"created_at"`
	UpdatedAt           DateFormattedTime `json:"updated_at"`
}

type ManutencaoUpdate struct {
	KmAtual             int64     `json:"km_atual"`
	Descricao           string    `json:"descricao"`
	ValorPago           float32   `json:"valor_pago"`
	VeiculoId           int64     `json:"veiculo_id"`
	UsuarioId           int64     `json:"usuario_id"`
	ManutencaoTipoId    int64     `json:"manutencao_tipo_id"`
	IsFinalizado        bool      `json:"is_finalizado"`
	IsAndamento         bool      `json:"is_andamento"`
	IsCancelado         bool      `json:"is_cancelado"`
	CanceladoEm         time.Time `json:"cancelado_em"`
	VeiculoIdTemporario int64     `json:"veiculo_id_temporario" `
}

type Manuntencoes struct {
	Id                  int64             `json:"id"`
	KmAtual             int64             `json:"km_atual"`
	Descricao           string            `json:"descricao"`
	ValorPago           float32           `json:"valor_pago"`
	CategoriaManutencao string            `json:"categoria_manutencao"`
	StatusManutencao    string            `json:"status_manutencao"`
	IdVeiculoManutencao int64             `json:"id_veiculo_manutencao"`
	VeiculoManutencao   string            `json:"veiculo_manutencao"`
	VeiculoSubstituto   string            `json:"veiculo_substituto"`
	IdVeiculoSubstituto int64             `json:"id_veiculo_substituto"`
	UsuarioId           int64             `json:"usuario_id"`
	Usuario             string            `json:"usuario"`
	CanceladoEm         DateFormattedTime `json:"cancelado_em"`
	RegistradoEm        DateFormattedTime `json:"registrado_em"`
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
