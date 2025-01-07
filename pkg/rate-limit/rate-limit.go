package ratelimit

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"wget/pkg"
	"wget/pkg/progress"
	"wget/pkg/utils"
)

func DownloadRateLimit(url, outputDir, rate string) error {
	file_path, _ := utils.GetFileInfo(url)
	outFile, err := os.Create(file_path.FileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	fmt.Printf("start at %s", utils.FormatDate(time.Now()))

	res, err := utils.CheckLink(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	infos, err := utils.GetFileInfo(url)
	if err != nil {
		return err
	}

	fmt.Printf("sending request, awaiting response... ")
	fmt.Printf("status %v\n", utils.HttpCode(infos.StatusCode))

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("URL %s returned status code %d", url, res.StatusCode)
	}
	fmt.Printf("content size: %s [~%s]\n", infos.ContentLength, utils.ToMega(infos.ContentLength))

	if *pkg.Output != "" {
		if !strings.HasPrefix(*pkg.NewPath, "./") {
			fmt.Printf("saving file to: %s\n", filepath.Join(filepath.Clean(*pkg.NewPath), *pkg.Output))
		} else {
			fmt.Printf("saving file to: ./%s\n", filepath.Join(filepath.Clean(*pkg.NewPath), *pkg.Output))
		}
	} else {
		if !strings.HasPrefix(*pkg.NewPath, "./") {
			fmt.Printf("saving file to: %s\n", filepath.Join(filepath.Clean(*pkg.NewPath), infos.FileName))
		} else {
			fmt.Printf("saving file to: ./%s\n", filepath.Join(filepath.Clean(*pkg.NewPath), infos.FileName))
		}
	}

	//Affichage des informations sur le fichier
	totalSize := res.ContentLength
	// if totalSize < 0 {
	// 	fmt.Println("Length: unspecified [text/html]\n")
	// } else {
	// 	fmt.Println("Content-Length: " + strconv.FormatInt(totalSize, 10) + " (" + utils.ConvertSize(totalSize) + ")" + "\n")
	// }

	debit, err := utils.ParseRateLimit(rate)
	if err != nil {
		fmt.Println("rate error:", err)
		return err
	}
	limiteread := NewRateLimitedReader(res.Body, debit)
	reader := progress.Progression(limiteread, totalSize)

	_, err = io.Copy(outFile, reader)
	if err != nil {
		fmt.Println("Ecriture:", err)
		return err
	}

	// fmt.Println("Saving file to: %s", file_path.FileName)
	fmt.Printf("\n\nfinished at: %v\n", utils.FormatDate(time.Now()))
	return nil
}
