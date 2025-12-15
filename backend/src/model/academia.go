package model

type Academia struct {
	ID        uint     `json:"id"`
	Nome      string   `json:"nome"`
	Endereco  string   `json:"endereco"`
	Telefone  string   `json:"telefone"`
	Preco     string   `json:"preco"`
	Descricao string   `json:"descricao"`
	Imagens   []Imagem `json:"imagens" gorm:"foreignKey:AcademiaID"`
	// Campo computado - não armazenado no banco, apenas para facilitar o frontend
	ImagemPrincipal string `json:"imagem_principal" gorm:"-"`
}
