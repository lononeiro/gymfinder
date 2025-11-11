package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/lononeiro/gymfinder/backend/src/model"
	"github.com/lononeiro/gymfinder/backend/src/repository"
	"github.com/lononeiro/gymfinder/backend/src/utils"
)

const UploadPath = "./uploads"
const MaxUploadSize = 10 << 20

func AdicionarAcademia(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize+512)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, "Arquivo muito grande ou erro no envio", http.StatusBadRequest)
		return
	}

	academia := model.Academia{
		Nome:     r.FormValue("nome"),
		Endereco: r.FormValue("endereco"),
		Telefone: r.FormValue("telefone"),
		Preco:    r.FormValue("preco"),
	}

	file, header, err := r.FormFile("imagem")
	filename := ""

	if err == nil {
		defer file.Close()

		buf := make([]byte, 512)
		n, _ := file.Read(buf)
		contentType := http.DetectContentType(buf[:n])
		types := map[string]bool{"image/jpeg": true, "image/png": true, "image/webp": true}
		if !types[contentType] {
			http.Error(w, "Formato de imagem não permitido", http.StatusBadRequest)
			return
		}

		file.Seek(0, io.SeekStart)
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext == "" {
			if contentType == "image/png" {
				ext = ".png"
			} else {
				ext = ".jpg"
			}
		}

		filename = uuid.New().String() + ext
		os.MkdirAll(UploadPath, os.ModePerm)

		dst, err := os.Create(filepath.Join(UploadPath, filename))
		if err != nil {
			http.Error(w, "Erro ao salvar imagem", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		io.Copy(dst, file)
		academia.Imagem = filename
	}

	created, err := repository.CreateAcademia(academia, model.Imagem{URL: filename})
	if err != nil {
		http.Error(w, "Erro ao criar academia: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Academia criada com sucesso",
		"academia": created,
	})
}

func ListarAcademias(w http.ResponseWriter, r *http.Request) {
	academias := repository.ListarAcademias()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(academias)
}

func EditarAcademias(w http.ResponseWriter, r *http.Request) {
	var academia model.Academia
	if err := json.NewDecoder(r.Body).Decode(&academia); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Erro ao decodificar o corpo da requisição: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	id, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}

	err := repository.EditarAcademias(id, academia)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Erro ao editar academia: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func ApagarAcademia(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}

	err := repository.ApagarAcademia(id)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Erro ao apagar academia: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
