package DB

import (
	"errors"
	"fmt"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/gorm"
)

func SelecionarUsuarioPorEmail(db *gorm.DB, email string) (model.Usuario, error) {
	var usuario model.Usuario
	err := db.Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return model.Usuario{}, err
	}
	return usuario, nil
}

func CreateUsuario(db *gorm.DB, usuario model.Usuario) error {

	_, err := SelecionarUsuarioPorEmail(db, usuario.Email)

	if err == nil {
		return fmt.Errorf("usuario com email %s já existe", usuario.Email)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err := db.Create(&usuario).Error
		if err != nil {
			return err
		}
		fmt.Println("Usuario criado com sucesso:", usuario.Nome)
		return nil
	}

	// Se for outro erro, retorna
	return err
}

// func SelecionarAcademiaPoriD(db *gorm.DB, id uint) (model.Academia, error) {
// 	var academia model.Academia
// 	err := db.First(&academia, id).Error
// 	if err != nil {
// 		return model.Academia{}, err
// 	}
// 	return academia, nil
// }

func ListarUsuarios(db *gorm.DB) ([]model.Usuario, error) {
	var usuarios []model.Usuario
	err := db.Find(&usuarios).Error
	if err != nil {
		fmt.Println("Erro ao listar usuarios:", err)
		return []model.Usuario{}, err
	}
	fmt.Printf("%d usuairos encontrados\n", len(usuarios))
	return usuarios, nil
}

func ApagarUsuario(db *gorm.DB, id uint) error {
	var usuairo model.Usuario
	err := db.First(&usuairo, id).Error
	if err != nil {
		return err
	}
	err = db.Delete(&usuairo).Error
	if err != nil {
		return err
	}
	fmt.Println("Usuario apagado com sucesso:", usuairo.Nome)
	return nil
}

//editar o usuario pelo id recebido
// func EditarUsuario(db *gorm.DB, id uint, usuario model.Usuario) error{
// 	var existingUsuario model.Academia
// 	err := db.First(&existingUsuario, id).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Atualiza os campos da academia existente com os novos valores
// 	existingUsuario.Nome = academia.Nome
// 	existingUsuario.Endereco = academia.Endereco
// 	existingUsuario.Telefone = academia.Telefone

// 	err = db.Save(&existingUsuario).Error
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Academia editada com sucesso:", existingUsuario.Nome)
// 	return nil
// }

// editar academia pelo id recebido
// func EditarAcademias(db *gorm.DB, id uint, academia model.Academia) error {
// 	var existingAcademia model.Academia
// 	err := db.First(&existingAcademia, id).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Atualiza os campos da academia existente com os novos valores
// 	existingAcademia.Nome = academia.Nome
// 	existingAcademia.Endereco = academia.Endereco
// 	existingAcademia.Telefone = academia.Telefone

// 	err = db.Save(&existingAcademia).Error
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Academia editada com sucesso:", existingAcademia.Nome)
// 	return nil
// }

func LoginUsuario(db *gorm.DB, email string, senha string) (model.Usuario, error) {
	var usuario model.Usuario
	err := db.Where("email = ? AND senha = ?", email, senha).First(&usuario).Error
	if err != nil {
		return model.Usuario{}, err
	}
	fmt.Println("Usuario logado com sucesso:", usuario.Nome)
	return usuario, nil
}
