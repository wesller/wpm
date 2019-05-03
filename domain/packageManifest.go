package domain

import (
	"log"
	"os"
	"path/filepath"

	"github.com/wesller/wpm/logger"
)

// PackageManifest Estrutura do arquivo de Manifesto
type PackageManifest struct {
	Name  string   `json:"name"`
	Hash  string   `json:"hash"`
	Files []string `json:"files"`
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

// Generate Gera o arquivo manifest de todos os arquivos da pasta
func Generate(dir string) {
	var files []string
	err := filepath.Walk(dir, visit(&files))
	if err != nil {
		logger.LogMsgFatal(err.Error())
	}

	for _, file := range files {
		logger.LogMsgInfo(file)
	}
}
