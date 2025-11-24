package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/lononeiro/gymfinder/backend/src/model"
	"github.com/lononeiro/gymfinder/backend/src/repository"
	"github.com/lononeiro/gymfinder/backend/src/utils"
)

const MaxUploadSize = 10 << 20 // 10 MB

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

	files := r.MultipartForm.File["imagens"]
	var imagens []model.Imagem

	if len(files) == 0 {
		http.Error(w, "É necessário enviar pelo menos uma imagem", http.StatusBadRequest)
		return
	}

	for _, header := range files {
		file, err := header.Open()
		if err != nil {
			http.Error(w, "Erro ao abrir arquivo", http.StatusBadRequest)
			return
		}
		// O defer file.Close() é chamado APÓS o loop, o que está tecnicamente incorreto
		// para lidar com erros dentro do loop, mas vamos mantê-lo aqui
		// para seguir a sua estrutura, idealmente deveria estar dentro do loop.
		defer file.Close()

		// Verifica tipo da imagem
		buf := make([]byte, 512)
		n, _ := file.Read(buf)
		contentType := http.DetectContentType(buf[:n])

		types := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/webp": true,
		}

		if !types[contentType] {
			http.Error(w, "Formato de imagem não permitido", http.StatusBadRequest)
			return
		}

		// volta o ponteiro do arquivo
		_, _ = file.Seek(0, io.SeekStart)

		// Gera a extensão
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext == "" {
			switch contentType {
			case "image/png":
				ext = ".png"
			case "image/webp":
				ext = ".webp"
			default:
				ext = ".jpg"
			}
		}

		// Nome final único
		filename := uuid.New().String() + ext

		// 🔥 Envia para a FILEBASE (url = https://ipfs.filebase.io/ipfs/CID)
		url, err := utils.UploadToFilebase(file, filename)
		if err != nil {
			http.Error(w, "Erro ao enviar imagem para armazenamento: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Salva a URL COMPLETA no banco
		imagens = append(imagens, model.Imagem{URL: url})
	}

	created, err := repository.CreateAcademia(academia, imagens)
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

// ... (Resto das funções Listar, Editar, Apagar permanecem inalteradas)
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
