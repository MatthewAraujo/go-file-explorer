package filesystem

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func ListAllFolders(root string) ([]string, error) {
	var folders []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Se houver um erro de permissão, simplesmente ignoramos e continuamos
		if err != nil {
			if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.EACCES {
				log.Printf("Permissão negada ao tentar acessar %s. Ignorando...", path)
				return nil // Retorna nil para ignorar o erro e continuar
			}
			return err
		}

		if info.IsDir() {
			folders = append(folders, path)
		}

		return nil
	})

	return folders, err
}
