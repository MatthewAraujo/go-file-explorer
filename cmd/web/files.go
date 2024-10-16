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
