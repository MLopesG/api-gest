package model

type Usuario struct {
	Id        int64             `json:"id"`
	Nome      string            `json:"nome"`
	IsAtivo   bool              `json:"is_ativo"`
	Senha     string            `json:"senha"`
	Cpf       string            `json:"cpf"`
	Email     string            `json:"email"`
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
