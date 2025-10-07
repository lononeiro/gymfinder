package model

type Imagem struct {
	ID         int   `gorm:"primaryKey;autoIncrement" json:"id"`
	URL        string `json:"url"`
	AcademiaID int   `gorm:"index;not null" json:"academia_id"`
}
