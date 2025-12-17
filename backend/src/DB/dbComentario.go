package DB

import (
	"fmt"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/gorm"
)

func SelecionarComentarioPorID(db *gorm.DB, id uint) (model.Comentario, error) {
	var comentario model.Comentario
	err := db.First(&comentario, id).Error
	if err != nil {
		return model.Comentario{}, fmt.Errorf("comentário com ID %d não encontrado: %w", id, err)
	}
	fmt.Println("Comentário encontrado:", comentario.Texto)
	return comentario, nil
}

func CreateComentario(db *gorm.DB, comentario model.Comentario) error {
	err := db.Create(&comentario).Error
	if err != nil {
		return err
	}
	fmt.Println("Comentário criado com sucesso:", comentario.Texto)
	return nil
}

func ApagarComentario(db *gorm.DB, id uint) error {
	var comentario model.Comentario
	err := db.First(&comentario, id).Error
	if err != nil {
		return fmt.Errorf("comentário com ID %d não encontrado: %w", id, err)
	}

	err = db.Delete(&comentario).Error
	if err != nil {
		return fmt.Errorf("erro ao apagar comentário: %w", err)
	}

	fmt.Println("Comentário apagado com sucesso:", comentario.Texto)
	return nil
}

func EditarComentario(db *gorm.DB, id uint, comentario model.Comentario) error {
	var existingComentario model.Comentario
	err := db.First(&existingComentario, id).Error
	if err != nil {
		return fmt.Errorf("comentário com ID %d não encontrado: %w", id, err)
	}
	// Atualiza os campos do comentário existente com os novos valores
	existingComentario.Texto = comentario.Texto
	existingComentario.UsuarioID = comentario.UsuarioID
	existingComentario.AcademiaID = comentario.AcademiaID

	err = db.Save(&existingComentario).Error
	if err != nil {
		return fmt.Errorf("erro ao editar comentário: %w", err)
	}

	fmt.Println("Comentário editado com sucesso:", existingComentario.Texto)
	return nil
}

func ListarComentariosPost(db *gorm.DB, academiaID uint) ([]model.Comentario, error) {
	var comentarios []model.Comentario
	err := db.Where("academia_id = ?", academiaID).Preload("Usuario").Find(&comentarios).Error
	if err != nil {
		fmt.Println("Erro ao listar comentários:", err)
		return nil, err
	}

	// Define o nome do usuário para cada comentário
	for i := range comentarios {
		if comentarios[i].Usuario.Nome != "" {
			comentarios[i].UsuarioNome = comentarios[i].Usuario.Nome
		}
	}

	fmt.Printf("%d comentários encontrados para a academia ID %d\n", len(comentarios), academiaID)
	return comentarios, nil
}

func SelecionarUsuarioComentario(db *gorm.DB, comentarioID uint) (model.Usuario, error) {
	var comentario model.Comentario
	err := db.Preload("Usuario").First(&comentario, comentarioID).Error
	if err != nil {
		return model.Usuario{}, fmt.Errorf("comentário com ID %d não encontrado: %w", comentarioID, err)
	}
	return comentario.Usuario, nil
}
