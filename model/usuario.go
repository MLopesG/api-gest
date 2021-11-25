package model

type Usuario struct {
	Id      int64  `json:"id"`
	Nome    string `json:"nome"`
	IsAtivo bool   `json:"is_ativo"`
	Senha   string `json:"senha"`
	Cpf     string `json:"cpf"`
	Email   string `json:"email"`
}
