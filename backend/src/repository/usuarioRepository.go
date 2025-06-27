package repository

import (
	"fmt"

	"github.com/lononeiro/gymfinder/backend/src/DB"
	"github.com/lononeiro/gymfinder/backend/src/model"
)

func CreateUsuario(usuairo model.Usuario) error {
	err := DB.CreateUsuario(DB.DataBase, usuairo)

	if err != nil {
		fmt.Println("Erro ao criar usuario:", err)
	}
	return err
}

// func ListarAcademias() []model.Academia {
// 	academias, err := DB.ListarAcademias(DB.DataBase)
// 	if err != nil {
// 		fmt.Println("Erro ao listar academias:", err)
// 		return []model.Academia{}
// 	}
// 	return academias
// }

// func EditarAcademias(id uint, academia model.Academia) error {
// 	existingAcademia, err := DB.SelecionarAcademiaPoriD(DB.DataBase, id)
// 	if err != nil {
// 		fmt.Println("Erro ao selecionar academia:", err)
// 		return err
// 	}

// 	// Atualiza os campos da academia existente com os novos valores
// 	existingAcademia.Nome = academia.Nome
// 	existingAcademia.Endereco = academia.Endereco
// 	existingAcademia.Telefone = academia.Telefone

// 	err = DB.EditarAcademias(DB.DataBase, id, existingAcademia)
// 	if err != nil {
// 		fmt.Println("Erro ao editar academia:", err)
// 		return err
// 	}
// 	return nil
// }

// func ApagarAcademia(id uint) error {
// 	err := DB.ApagarAcademia(DB.DataBase, id)
// 	if err != nil {
// 		fmt.Println("Erro ao apagar academia:", err)
// 	}
// 	return err
// }

func LoginUsuario(email string, senha string) (model.Usuario, error) {
	usuario, err := DB.LoginUsuario(DB.DataBase, email, senha)
	if err != nil {
		fmt.Println("Erro ao fazer login:", err)
		return model.Usuario{}, err
	}
	return usuario, nil
}
