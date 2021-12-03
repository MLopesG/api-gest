package model

import "time"

type ManutencaoPrevia struct {
	Id                 int64             `json:"id"`
	KmPrevisao         int64             `json:"km_previsao"`
	DataPrevisao       time.Time         `json:"data_previsao"`
	UltimaManutencaoId int64             `json:"ultima_manutencao_id"`
	IsAprovado         bool              `json:"is_aprovado"`
	IsCancelado        bool              `json:"is_cancelado"`
	CreatedAt          DateFormattedTime `json:"created_at"`
	UpdatedAt          DateFormattedTime `json:"updated_at"`
}

type ManutencaoPrevisaoDetalhada struct {
	Id                        int64             `json:"id"`
	KmPrevisao                int64             `json:"km_previsao"`
	DataPrevisao              int64             `json:"data_previsao"`
	StatusPrevia              string            `json:"status_previsao"`
	KmAtual                   int64             `json:"km_atual"`
	Descricao                 string            `json:"descricao"`
	ValorPago                 float32           `json:"valor_pago"`
	CategoriaManutencao       string            `json:"categoria_manutencao"`
	StatusManutencao          string            `json:"status_manutencao"`
	IdVeiculoManutencao       int64             `json:"id_veiculo_manutencao"`
	VeiculoManutencao         string            `json:"veiculo_manutencao"`
	VeiculoSubstituto         string            `json:"veiculo_substituto"`
	IdVeiculoSubstituto       int64             `json:"id_veiculo_substituto"`
	UsuarioId                 int64             `json:"usuario_id"`
	Usuario                   string            `json:"usuario"`
	ManutencaoRealizada       DateFormattedTime `json:"manutencao_realizado"`
	PrevisaoAbertaEm          DateFormattedTime `json:"previsao_aberta_em"`
	UltimaAtualizacaoRegistro DateFormattedTime `json:"ultima_atualizacao_registro"`
}

type ManutencaoPrevisaoList struct {
	Id                        int64             `json:"id"`
	KmPrevisao                int64             `json:"km_previsao"`
	DataPrevisao              int64             `json:"data_previsao"`
	CategoriaManutencao       string            `json:"categoria_manutencao"`
	StatusPrevia              string            `json:"status_previsao"`
	Descricao                 string            `json:"descricao"`
	VeiculoManutencao         string            `json:"veiculo_manutencao"`
	PrevisaoAbertaEm          DateFormattedTime `json:"previsao_aberta_em"`
	UltimaAtualizacaoRegistro DateFormattedTime `json:"ultima_atualizacao_registro"`
}
