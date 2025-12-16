package DB

import (
	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/gorm"
)

const UploadPath = "./uploads"

func CreateAcademia(db *gorm.DB, academia model.Academia, imagens []model.Imagem) (model.Academia, error) {
	tx := db.Begin()

	if err := tx.Create(&academia).Error; err != nil {
		tx.Rollback()
		return academia, err
	}

	for i := range imagens {
		imagens[i].AcademiaID = academia.ID

		if err := tx.Create(&imagens[i]).Error; err != nil {
			tx.Rollback()
			return academia, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return academia, err
	}

	return academia, nil
}

func SelecionarAcademiaPoriD(db *gorm.DB, id uint) (model.Academia, error) {
	var academia model.Academia
	err := db.Preload("Imagens").First(&academia, id).Error
	return academia, err
}

// ListarAcademias retorna academias com imagens e define a imagem principal
func ListarAcademias(db *gorm.DB) ([]model.Academia, error) {
	var academias []model.Academia

	err := db.Preload("Imagens").Find(&academias).Error
	if err != nil {
		return nil, err
	}

	// Define a imagem principal para cada academia (primeira imagem do array)
	for i := range academias {
		if len(academias[i].Imagens) > 0 {
			academias[i].ImagemPrincipal = academias[i].Imagens[0].URL
		}
	}

	return academias, nil
}

// --------------------------------------------------
// APAGAR ACADEMIA
// --------------------------------------------------
func ApagarAcademia(db *gorm.DB, id uint) error {

	var academia model.Academia
	err := db.Preload("Imagens").First(&academia, id).Error
	if err != nil {
		return err
	}

	// IPFS não apaga arquivo local → só remove do banco
	err = db.Delete(&academia).Error
	return err
}
func EditarAcademias(db *gorm.DB, id uint, academia model.Academia) error {
	return db.Save(&academia).Error
}
