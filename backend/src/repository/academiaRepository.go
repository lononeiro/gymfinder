package repository

import (
	"fmt"

	"github.com/lononeiro/gymfinder/backend/src/DB"
	"github.com/lononeiro/gymfinder/backend/src/model"
)

func CreateAcademia(academia model.Academia, imagens []model.Imagem) (model.Academia, error) {
	created, err := DB.CreateAcademia(DB.DataBase, academia, imagens)
	if err != nil {
		fmt.Println("Erro ao criar academia:", err)
	}
	return created, err
}

func ListarAcademias() []model.Academia {
	academias, err := DB.ListarAcademias(DB.DataBase)
	if err != nil {
		fmt.Println("Erro ao listar academias:", err)
		return []model.Academia{}
	}
	return academias
}

func EditarAcademias(id uint, academia model.Academia) error {
	existingAcademia, err := DB.SelecionarAcademiaPoriD(DB.DataBase, id)
	if err != nil {
		return err
	}

	existingAcademia.Nome = academia.Nome
	existingAcademia.Endereco = academia.Endereco
	existingAcademia.Telefone = academia.Telefone
	existingAcademia.Preco = academia.Preco

	err = DB.EditarAcademias(DB.DataBase, id, existingAcademia)
	return err
}

func ApagarAcademia(id uint) error {
	err := DB.ApagarAcademia(DB.DataBase, id)
	return err
}

func ObterAcademiaPorID(id uint) (model.Academia, error) {
	academia, err := DB.SelecionarAcademiaPoriD(DB.DataBase, id)
	return academia, err
}
