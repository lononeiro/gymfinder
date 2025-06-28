package repository

import (
	"fmt"

	"github.com/lononeiro/gymfinder/backend/src/DB"
	"github.com/lononeiro/gymfinder/backend/src/model"
)

func SelecionarComentarioPorID(id uint) (model.Comentario, error) {
	comentario, err := DB.SelecionarComentarioPorID(DB.DataBase, id)
	if err != nil {
		fmt.Println("Erro ao selecionar comentário:", err)
		return model.Comentario{}, err
	}
	return comentario, nil
}

func CriarComentario(Comentario model.Comentario) error {
	err := DB.CreateComentario(DB.DataBase, Comentario)

	if err != nil {
		fmt.Println("Erro ao criar comentário:", err)
	}
	return err
}

func ApagarComentario(id uint) error {
	err := DB.ApagarComentario(DB.DataBase, id)

	if err != nil {
		fmt.Println("Erro ao apagar comentário:", err)
		return err
	}
	return err
}

func EditarComentario(id uint, comentario model.Comentario) error {
	err := DB.EditarComentario(DB.DataBase, id, comentario)

	if err != nil {
		fmt.Println("Erro ao editar comentário:", err)
		return err
	}
	return nil
}

func ListarComentariosPost(academiaID uint) ([]model.Comentario, error) {
	comentarios, err := DB.ListarComentariosPost(DB.DataBase, academiaID)
	if err != nil {
		fmt.Println("Erro ao listar comentários:", err)
		return []model.Comentario{}, err
	}
	return comentarios, nil
}
