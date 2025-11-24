package DB

import (
	"os"
	"path/filepath"

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

	// ⚠️ ATENÇÃO: NENHUMA LÓGICA DE MODIFICAÇÃO DE URL É NECESSÁRIA AQUI.
	// As URLs de IPFS devem ser retornadas EXATAMENTE como estão salvas no banco.

	return academias, nil
}
func ApagarAcademia(db *gorm.DB, id uint) error {
	var academia model.Academia
	err := db.Preload("Imagens").First(&academia, id).Error
	if err != nil {
		return err
	}

	// NOTA: Esta remoção de arquivo só funcionaria para uploads locais,
	// mas não afetará o Filebase/IPFS. Manter a lógica de remoção de arquivo
	// é perigosa aqui, pois o URL no banco agora é um CID/URL IPFS.
	// Vou manter seu código original aqui, mas note que não apaga no Filebase.
	for _, img := range academia.Imagens {
		os.Remove(filepath.Join(UploadPath, img.URL))
	}

	err = db.Delete(&academia).Error
	return err
}

func EditarAcademias(db *gorm.DB, id uint, academia model.Academia) error {
	return db.Save(&academia).Error
}
