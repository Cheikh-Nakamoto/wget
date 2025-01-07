package utils

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"wget/pkg"
)

type FileInfo struct {
	StatusCode    int
	ContentType   string
	ContentLength string
	FileName      string
}

func GetFileInfo(url string) (*FileInfo, error) {
	response, err := http.Head(url)
	if err != nil {
		return nil, fmt.Errorf("impossible de récupérer les informations du fichier: %v", err)
	}
	defer response.Body.Close()

	return &FileInfo{
		StatusCode:    response.StatusCode,
		ContentType:   response.Header.Get("Content-Type"),
		ContentLength: fmt.Sprintf("%v", response.ContentLength),
		FileName:      filepath.Base(url),
	}, nil
}

func CheckFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func CheckFileExtension(filename string) bool {
	return strings.ToLower(filepath.Ext(filename)) == ".txt"
}

func ReadFileContent(filename string) ([]string, error) {
	if !CheckFileExist(filename) {
		return nil, fmt.Errorf("le fichier %s n'existe pas", filename)
	}

	if !CheckFileExtension(filename) {
		return nil, fmt.Errorf("le fichier %s n'est pas un fichier texte", filename)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := bufio.NewScanner(strings.NewReader(string(content)))
	var urls []string
	for lines.Scan() {
		line := strings.TrimSpace(lines.Text())

		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "/") || strings.HasPrefix(line, "$") {
			continue
		}

		if !CheckURL(line) {
			return nil, fmt.Errorf("le lien %s n'est pas valide", line)
		}
		urls = append(urls, line)
	}
	return urls, nil
}

func RenameFile(oldName, newName string) error {
	dir, err := AbsolutePath(*pkg.NewPath)
	if err != nil {
		return err
	}

	return os.Rename(filepath.Join(dir, oldName), filepath.Join(dir, newName))
	//return os.Rename(fmt.Sprintf("%v%v", dir, oldName), fmt.Sprintf("%v%v", dir, newName))
	//return os.Rename(oldPath, newPath)
}

func CreateNewFile(filename, path string) error {
	output := filepath.Join(path, filename)
	if !CheckFileExist(output) {
		file, err := os.Create(output)
		if err != nil {
			return fmt.Errorf("error while creating file %s: %v", output, err)
		}
		file.Close()
	}
	return nil
}

func Sizes(files []string) []int {
	var result []int
	for _, url := range files {
		infos, err := GetFileInfo(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		value, err := strconv.Atoi(infos.ContentLength)
		if err != nil {
			fmt.Println("cannot convert", infos.ContentLength)
		}

		result = append(result, value)
	}
	sort.Ints(result)

	return result
}

func Display(tab []int) string {
	var result string
	result = "["
	for i, val := range tab {
		if i != 0 {
			result += ", "
		}
		result += strconv.Itoa(val)
	}
	result += "]"

	return result
}

func WriteLog( message string) error {
	// Ouvrir (ou créer) le fichier de log en mode ajout
	file, err := os.OpenFile("wget-log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	// Écrire le message dans le fichier
	_, err = file.WriteString(message + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to log file: %v", err)
	}

	return nil
}