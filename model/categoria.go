package model

type Categoria struct {
	Id                    int               `json:"id" `
	Nome                  string            `json:"nome"`
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
