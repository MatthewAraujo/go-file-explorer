package server

import (
	"net/http"

	"github.com/MatthewAraujo/go-file-explorer/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)
	mux.HandleFunc("/files", web.ListFilesHandler)
	mux.HandleFunc("/subdirectories", web.ListSubDirectoriesHandler)

	return mux
}
