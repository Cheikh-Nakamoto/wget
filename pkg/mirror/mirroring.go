package mirror

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"wget/pkg"
)

// FilterLinks permet de filtrer les liens selon les options (-R: fichier ou -X: repertoires) à rejeter ou exclure
// links: liste de liens à filtrer
// Retourne une liste de liens filtrée
// Cette fonction reste encore à voir concernant des types d'entrée
func AllowedSave(filename string) bool {
	flag.Parse()
	// Exclude folders
	var X []string
	if *pkg.Exclude != "" {
		X = strings.Split(*pkg.Exclude, ",")
	} else if *pkg.XLong != "" {
		X = strings.Split(*pkg.XLong, ",")
	}

	excluded_folder := false
	for _, excl := range X {
		if strings.HasPrefix(excl, "/") {
			excl = excl[1:]
		}
		sanitized_folder_name := strings.TrimPrefix(strings.ReplaceAll(*pkg.NewPath+"/"+excl, "//", "/"), "./")
		if len(excl) > 0 && strings.HasPrefix(filename, sanitized_folder_name) {
			excluded_folder = true
		}
	}

	// Reject file types
	var rejections []string
	if *pkg.Reject != "" {
		rejections = strings.Split(*pkg.Reject, ",")
	} else if *pkg.RLong != "" {
		rejections = strings.Split(*pkg.RLong, ",")
	}
	// rejections := strings.Split(*pkg.Reject, ",")
	rejected_type := false
	for i := range rejections {
		extension := "." + rejections[i]
		if strings.HasSuffix(filename, extension) {
			rejected_type = true
		}
	}

	return !excluded_folder && !rejected_type
}

// ConvertLinks permet de convertir les liens pour une utilisation hors ligne
// file: fichier contenant les liens à convertir
// Retourne une erreur si la conversion échoue
// Cette fonction reste encore à voir concernant des types d'entrée
func ConvertLinks(content string) string {
	// Créez une expression régulière pour correspondre à l'URL complète
	re1 := regexp.MustCompile(regexp.QuoteMeta(RootURL))

	if re1.MatchString(content) {
		content = re1.ReplaceAllString(content, "/")

		// Remplacez les doubles slashes "//" par un simple slash "/"
		// en évitant de modifier les "://" présents dans les protocoles (http://, https://)
		re2 := regexp.MustCompile(`([^:])/+`)
		content = re2.ReplaceAllString(content, "$1/")
	}

	return content
}

func MirrorSite(url, outputPath string) error {
	dirPath := filepath.Join("./", outputPath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Erreur lors de la création du répertoire : %v\n", err)
		return err
	}
	visited := make(map[string]bool)
	DownloadLinks(url, dirPath, url, 5, 0, visited)
	return nil
}
