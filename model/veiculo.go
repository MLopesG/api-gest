package model

type Veiculo struct {
	Id          int               `json:"id"`
	Placa       string            `json:"placa"`
	Descricao   string            `json:"descricao"`
	CategoriaId int               `json:"categoria_id"`
	IsServico   bool              `json:"is_servico"`
	CreatedAt   DateFormattedTime `json:"created_at"`
	UpdatedAt   DateFormattedTime `json:"updated_at"`
}

type VeiculoUpdate struct {
	Placa       string `json:"placa"`
	Descricao   string `json:"descricao"`
	CategoriaId int    `json:"categoria_id"`
	IsServico   bool   `json:"is_servico"`
}
