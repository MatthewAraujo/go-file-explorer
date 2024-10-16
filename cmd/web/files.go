package web

import (
	"log"
	"net/http"

	"github.com/MatthewAraujo/go-file-explorer/internal/filesystem"
)

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	directories, err := filesystem.ListAll("/home")
	if err != nil {
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		return
	}

	fakeStorage := filesystem.NewFakeStorage()
	dic := fakeStorage.DisplayTree(directories.Name)

	fileComponent := FileSearched()
	err = fileComponent.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering in ListFilesHandler: %e", err)
	}

	component := FilesList(dic, "home")
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering in ListFilesHandler: %e", err)
	}
}

func ListSubDirectoriesHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("directory")
	if path == "" {
		path = "/"
	}

	fakeStorage := filesystem.NewFakeStorage()
	directories := fakeStorage.DisplayTree(path)

	if directories == nil {
		http.Error(w, "directory not found", http.StatusNotFound)
		return
	}

	component := FilesList(directories, path)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering filesList %e", err)
	}

}

func SearchFileHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to process form data", http.StatusInternalServerError)
		return
	}

	file := r.FormValue("file")
	if file == "" {
		http.Error(w, "File not provided", http.StatusBadRequest)
		return
	}

	fakeStorage := filesystem.NewFakeStorage()
	files, err := fakeStorage.SearchFile(file)
	if err != nil {
		component := FileSearchedResult(files)
		err := component.Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering result", http.StatusInternalServerError)
		}
		return
	}

	// Renderiza os arquivos encontrados
	component := FileSearchedResult(files)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering FileSearchedResult %e", err)
	}
}
