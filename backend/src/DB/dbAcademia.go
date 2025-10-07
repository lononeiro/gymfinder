package DB

import (
	"fmt"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/gorm"
)


func CreateAcademia(db *gorm.DB, academia model.Academia, imagem model.Imagem)  (model.Academia, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&academia).Error; err != nil {
			return err
		}
		fmt.Println("Academia criada com sucesso:", academia.Nome, academia.ID)
		
		if imagem.URL != "" {
			imagem.AcademiaID = academia.ID

			if err := tx.Create(&imagem).Error; err != nil {
				return err
			}
			fmt.Println("Imagem criada com sucesso para a academia:", imagem.URL)
	}
		return nil
	})
	return academia, err
}

// 	err := db.Create(&academia).Error
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Academia criada com sucesso:", academia.Nome)
	
// 	imagem.AcademiaID = academia.ID
	
// 	err = db.Create(&imagem).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil


func SelecionarAcademiaPoriD(db *gorm.DB, id uint) (model.Academia, error) {
	var academia model.Academia
	err := db.First(&academia, id).Error
	if err != nil {
		return model.Academia{}, err
	}
	return academia, nil
}

func ListarAcademias(db *gorm.DB) ([]model.Academia, error) {
	var academias []model.Academia
	err := db.Find(&academias).Error
	if err != nil {
		fmt.Println("Erro ao listar academias:", err)
		return []model.Academia{}, err
	}
	fmt.Printf("%d academias encontradas\n", len(academias))
	return academias, nil
}

func ApagarAcademia(db *gorm.DB, id uint) error {
	var academia model.Academia
	err := db.First(&academia, id).Error
	if err != nil {
		return err
	}
	err = db.Delete(&academia).Error
	if err != nil {
		return err
	}
	fmt.Println("Academia apagada com sucesso:", academia.Nome)
	return nil
}

// editar academia pelo id recebido
func EditarAcademias(db *gorm.DB, id uint, academia model.Academia) error {
	var existingAcademia model.Academia
	err := db.First(&existingAcademia, id).Error
	if err != nil {
		return err
	}

	// Atualiza os campos da academia existente com os novos valores
	existingAcademia.Nome = academia.Nome
	existingAcademia.Endereco = academia.Endereco
	existingAcademia.Telefone = academia.Telefone

	err = db.Save(&existingAcademia).Error
	if err != nil {
		return err
	}
	fmt.Println("Academia editada com sucesso:", existingAcademia.Nome)
	return nil
}
