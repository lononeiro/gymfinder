package repository

import (
	"fmt"

	"github.com/lononeiro/gymfinder/backend/src/DB"
	"github.com/lononeiro/gymfinder/backend/src/model"
)

func CreateAcademia(academia model.Academia) error {
	err := DB.CreateAcademia(DB.DB, academia)

	if err != nil {
		fmt.Println("Erro ao criar academia:", err)
	}
	return err
}

func ListarAcademias() []model.Academia {
	academias, err := DB.ListarAcademias(DB.DB)
	if err != nil {
		fmt.Println("Erro ao listar academias:", err)
		return []model.Academia{}
	}
	return academias
}

func EditarAcademias(id uint, academia model.Academia) error {
	existingAcademia, err := DB.SelecionarAcademiaPoriD(DB.DB, id)
	if err != nil {
		fmt.Println("Erro ao selecionar academia:", err)
		return err
	}

	// Atualiza os campos da academia existente com os novos valores
	existingAcademia.Nome = academia.Nome
	existingAcademia.Endereco = academia.Endereco
	existingAcademia.Telefone = academia.Telefone

	err = DB.EditarAcademias(DB.DB, id, existingAcademia)
	if err != nil {
		fmt.Println("Erro ao editar academia:", err)
		return err
	}
	return nil
}

func ApagarAcademia(id uint) error {
	err := DB.ApagarAcademia(DB.DB, id)
	if err != nil {
		fmt.Println("Erro ao apagar academia:", err)
	}
	return err
}
