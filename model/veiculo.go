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

type VeiculoCategoria struct {
	Id                    int    `json:"id"`
	Placa                 string `json:"placa"`
	CategoriaId           int    `json:"categoria_id"`
	NomeCategoria         string `json:"nome_categoria"`
	Descricao             string `json:"descricao"`
	IsServico             bool   `json:"is_servico"`
	IsCategoriaProduto    bool   `json:"is_categoria_produto"`
	IsCategoriaManutencao bool   `json:"is_categoria_manutencao"`
	IsCategoriaVeiculo    bool   `json:"is_categoria_veiculo"`
}

type VeiculoUpdate struct {
	Placa       string `json:"placa"`
	Descricao   string `json:"descricao"`
	CategoriaId int    `json:"categoria_id"`
	IsServico   bool   `json:"is_servico"`
}
