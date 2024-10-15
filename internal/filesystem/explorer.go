package filesystem

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

// ListFirstLevelFolders lista os diretórios de primeiro nível no caminho especificado.
func ListFirstLevelFolders(root string) ([]string, error) {
	var folders []string

	dir, err := os.Open(root)
	if err != nil {
		log.Printf("Erro ao abrir diretório %s: %s", root, err)
		return nil, err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(0)
	if err != nil {
		log.Printf("Erro ao ler nomes de diretórios em %s: %s", root, err)
		return nil, err
	}

	for _, name := range names {
		if name == "." || name == ".." { // Ignora diretórios especiais
			continue
		}

		path := filepath.Join(root, name)
		info, err := os.Stat(path)
		if err != nil {
			pathErr, ok := err.(*os.PathError)
			if ok && pathErr.Err == syscall.EACCES {
				log.Printf("Permissão negada ao tentar acessar %s. Ignorando...", path)
				continue
			}
			log.Printf("Erro ao obter informações do diretório %s: %s", path, err)
			continue
		}

		if info.IsDir() {
			folders = append(folders, name) // Retorna apenas o nome do diretório
		}
	}

	return folders, nil
}
