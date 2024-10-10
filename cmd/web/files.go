package web

import (
	"log"
	"net/http"

	"github.com/MatthewAraujo/go-file-explorer/internal/filesystem"
)

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	directories, err := filesystem.ListAllFolders("/home")
	if err != nil {
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		return
	}

	// Renderize o template FilesList, passando a lista de diretórios
	component := FilesList(directories) // Chama o template com a lista
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering in ListFilesHandler: %e", err)
	}
}
