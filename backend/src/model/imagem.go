package model

type Imagem struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	URL        string   `json:"url"`
	AcademiaID uint     `json:"academia_id"`
	Academia   Academia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
