package model

type Academia struct {
	ID       uint     `json:"id"`
	Nome     string   `json:"nome"`
	Endereco string   `json:"endereco"`
	Telefone string   `json:"telefone"`
	Preco    string   `json:"preco"`
	Imagem   string   `json:"imagem"`
	Imagens  []Imagem `json:"imagens" gorm:"foreignKey:AcademiaID"`
}
