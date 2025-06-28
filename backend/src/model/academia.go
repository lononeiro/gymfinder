package model

type Academia struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Telefone string `json:"telefone"`
	Preco    string `json:"preco"`
}
