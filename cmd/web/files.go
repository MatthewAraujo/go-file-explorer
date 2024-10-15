package web

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/MatthewAraujo/go-file-explorer/internal/filesystem"
)

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	directories, err := filesystem.ListFirstLevelFolders("/home")
	if err != nil {
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		return
	}

	firstLevelDirs := extractFirstLevelDirectories(directories)

	component := FilesList(firstLevelDirs)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering in ListFilesHandler: %e", err)
	}
}

func extractFirstLevelDirectories(directories []string) []string {
	firstLevelDirs := make(map[string]struct{})

	for _, dir := range directories {
		dirName := strings.TrimSuffix(filepath.Base(dir), "/") // Remove a barra no final
		if dirName != "" {
			firstLevelDirs[dirName] = struct{}{}
		}
	}

	// Converter o mapa em um slice para exibição
	var displayDirectories []string
	for dir := range firstLevelDirs {
		displayDirectories = append(displayDirectories, dir)
	}
	return displayDirectories
}

func ListSubDirectoriesHandler(w http.ResponseWriter, r *http.Request) {
	directory := r.URL.Path[len("/subdirectories/"):] // Extrai o subdiretório

	// Garante que o caminho seja absoluto
	fullPath := filepath.Join("/home", directory)
	log.Printf("Recebido pedido para o diretório: %s", fullPath)

	directories, err := filesystem.ListFirstLevelFolders(fullPath)
	if err != nil {
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		log.Printf("Erro ao listar diretórios: %s", err) // Log do erro
		return
	}

	firstLevelDirs := extractFirstLevelDirectories(directories)

	log.Printf("Subdiretórios encontrados: %v", firstLevelDirs) // Log dos diretórios encontrados

	component := FilesList(firstLevelDirs)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Erro ao renderizar em ListSubDirectoriesHandler: %e", err)
	}
}
