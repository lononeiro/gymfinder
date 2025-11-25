package DB

import (
	"fmt"
	"os"
	"strings"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/gorm"
)

const UploadPath = "./uploads"

func CreateAcademia(db *gorm.DB, academia model.Academia, imagens []model.Imagem) (model.Academia, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&academia).Error; err != nil {
			return err
		}

		for i := range imagens {
			imagens[i].AcademiaID = academia.ID
		}

		if len(imagens) > 0 {
			if err := tx.Create(&imagens).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return academia, err
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

	// 🔥 CORREÇÃO: Ajuste automático de imagens antigas
	bucket := strings.TrimSpace(os.Getenv("FILEBASE_BUCKET"))
	apiBase := strings.TrimSpace(os.Getenv("API_BASE_URL")) // ex: https://gymfinder-1.onrender.com

	for ai := range academias {
		for ii := range academias[ai].Imagens {

			img := &academias[ai].Imagens[ii]

			// já é uma URL completa → nada a corrigir
			if strings.HasPrefix(img.URL, "http://") || strings.HasPrefix(img.URL, "https://") {
				continue
			}

			// tentar montar URL Filebase/S3
			if bucket != "" {
				img.URL = fmt.Sprintf("https://s3.filebase.com/%s/%s", bucket, img.URL)
				continue
			}

			// fallback: URL pública do backend
			if apiBase != "" {
				img.URL = fmt.Sprintf("%s/uploads/%s", strings.TrimRight(apiBase, "/"), img.URL)
				continue
			}

			// último caso: URL relativa
			img.URL = fmt.Sprintf("/uploads/%s", img.URL)
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
