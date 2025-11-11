package DB

import (
	"os"
	"path/filepath"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/gorm"
)

const UploadPath = "./uploads"

func CreateAcademia(db *gorm.DB, academia model.Academia, imagem model.Imagem) (model.Academia, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&academia).Error; err != nil {
			return err
		}

		if imagem.URL != "" {
			imagem.AcademiaID = academia.ID
			if err := tx.Create(&imagem).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return academia, err
}

func SelecionarAcademiaPoriD(db *gorm.DB, id uint) (model.Academia, error) {
	var academia model.Academia
	err := db.First(&academia, id).Error
	return academia, err
}

func ListarAcademias(db *gorm.DB) ([]model.Academia, error) {
	var academias []model.Academia
	err := db.Find(&academias).Error
	return academias, err
}

func ApagarAcademia(db *gorm.DB, id uint) error {
	var academia model.Academia
	err := db.First(&academia, id).Error
	if err != nil {
		return err
	}

	if academia.Imagem != "" {
		os.Remove(filepath.Join(UploadPath, academia.Imagem))
	}

	err = db.Delete(&academia).Error
	return err
}

func EditarAcademias(db *gorm.DB, id uint, academia model.Academia) error {
	return db.Save(&academia).Error
}
