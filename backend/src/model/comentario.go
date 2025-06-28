package model

type Comentario struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	AcademiaID uint     `gorm:"foreignKey:AcademiaID" json:"academia_id"`
	UsuarioID  uint     `gorm:"foreignKey:UsuarioID" json:"usuario_id"`
	Texto      string   `gorm:"type:text;not null" json:"texto"`
	Academia   Academia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Usuario    Usuario  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
