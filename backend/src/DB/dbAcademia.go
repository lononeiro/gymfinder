package DB

import (
	"strings"

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

// ListarAcademias AGORA NÃO PREFIXA MAIS A URL, POIS ELA JÁ DEVE ESTAR COMPLETA NO BANCO
// DB/DB.go

func ListarAcademias(db *gorm.DB) ([]model.Academia, error) {
	var academias []model.Academia

	err := db.Preload("Imagens").Find(&academias).Error
	if err != nil {
		return nil, err
	}

	// -------------- CONFIGURAÇÃO DO SEU GATEWAY -----------------
	// Seu gateway personalizado do Filebase
	const gateway = "https://future-coffee-galliform.myfilebase.com/"
	// ----------------------------------------------------------------

	for ai := range academias {
		for ii := range academias[ai].Imagens {

			img := &academias[ai].Imagens[ii]

			if img.URL == "" {
				continue
			}

			// Se já começa com http, não mexe
			if strings.HasPrefix(img.URL, "http://") || strings.HasPrefix(img.URL, "https://") {
				continue
			}

			// Se veio apenas um CID ou nome, prefixa com seu gateway personalizado
			img.URL = gateway + strings.TrimPrefix(img.URL, "/")
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
