package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"wget/pkg/progress"
	"wget/pkg/utils"
)

func Download(background bool, url, dir string) error {
	newDir, err := utils.AbsolutePath(dir)
	if err != nil {
		return fmt.Errorf("failed to expand tilde in directory path: %v", err)
	}

	response, err := utils.CheckLink(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("URL %s returned status code %d", url, response.StatusCode)
	}

	totalSize := response.ContentLength
	reader := progress.Progression(response.Body, totalSize)

	filaName := filepath.Base(url)
	outputPath := filepath.Join(newDir, filepath.Base(filaName))
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier %s: %v", outputPath, err)
	}
	defer outFile.Close()

	if background {
		_, err = io.Copy(outFile, response.Body)
	} else {
		_, err = io.Copy(outFile, reader)
	}
	if err != nil {
		return fmt.Errorf("erreur lors de la copie du contenu de %s: %v", url, err)
	}
	return nil
}
