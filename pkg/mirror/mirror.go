package mirror

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"wget/pkg"
	"wget/pkg/progress"
	"wget/pkg/utils"

	"golang.org/x/net/html"
)

type Link struct {
	Href string `json:"href"`
}

type Page struct {
	Response io.Reader
	size     int
	Links    []Link `json:"links"`
}

var RootURL = ""

func DownloadPage(url string) (*Page, error) {

	var page Page
	// Créer un client HTTP
	client := &http.Client{}

	// Créer une requête HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête :", err)
		return &page, err
	}

	// Ajouter des en-têtes courants d'un navigateur
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("HTTP request sent, awaiting response... " + resp.Status + "\n")
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %s", resp.Status)
	}
	//Affichage des informations sur le fichier
	totalSize := resp.ContentLength
	if totalSize < 0 {
		totalSize *= -1
		fmt.Println("Length: unspecified [text/html]")
	} else {
		fmt.Println("Content-Length: " + strconv.FormatInt(totalSize, 10) + " (" + utils.ConvertSize(totalSize) + ")" + "\n")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error: ", err)
		return nil, err
	}

	// Utiliser une copie du buffer pour extraire les liens
	reader := strings.NewReader(string(body))
	response := strings.NewReader(string(body))
	links, err := extractLinks(reader)
	if err != nil {
		fmt.Println("Erreur lors de l'extraction des liens :", err)
		return nil, err
	}
	page.Links = links
	page.Response = response
	page.size = int(totalSize)
	return &page, nil
}

func DownloadLinks(url, dirpath, linkbase string, maxDepth int, currentDepth int, visited map[string]bool) {
	if currentDepth > maxDepth || visited[url] {
		// fmt.Println("Download maxdepth reached 5 or link visited")
		return
	}

	pages, err := DownloadPage(url)
	if err != nil {
		// fmt.Println("DownloadPage error: ", err)
		return
	}
	visited[url] = true
	SavingFile(*pages, url, linkbase, dirpath)
	for _, link := range pages.Links {
		newUrl, err := resolveURL(url, link.Href)
		if err != nil {
			// fmt.Println("Erreur lors de la résolution de l'URL :", err)
			continue
		}

		// Vérifier que l'URL n'est pas encore visitée ou est hors site
		if visited[newUrl] {
			// fmt.Println("Lien déjà visité :", newUrl)
			continue
		}
		Link_log(link)
		DownloadLinks(newUrl, dirpath, linkbase, maxDepth, currentDepth+1, visited)
	}
}
func SavingFile(page Page, url, linkbase, dirpath string) error {
	filename, err := GetFileNameFromLink(url, linkbase, dirpath)
	if err != nil {
		fmt.Println("Get filename error: ", err)
		return err
	}

	re := regexp.MustCompile(`(https?:|#)`)
	
	if AllowedSave(filename) && !re.MatchString(filename) {
		// Créer les répertoires nécessaires
		err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			// fmt.Println("Erreur lors de la création des répertoires: ", err)
			return err
		}

		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Creation file error: ", err)
			return err
		}
		defer file.Close()

		// Convert reader to string
		buf := new(bytes.Buffer)
		buf.ReadFrom(page.Response)
		content := buf.String()
		// Check if file is a web page
		re := regexp.MustCompile(`^[\w/]+\.(html|php|css|js)$`)
		if *pkg.ConvertLinks && re.MatchString(filename) {
			content = ConvertLinks(content)
		}
		newread := strings.NewReader(content)
		reader := progress.Progression(newread, int64(page.size))
		_, err = io.Copy(file, reader)

		if err != nil {
			fmt.Println("Savefile error: ", err)
			return err
		}
		fmt.Println("File saved:", filename)
	}
	return nil
}

func GetFileNameFromLink(link, linkbase, base string) (string, error) {
	// j'enleve la base dans le lien actuelle
	filename := strings.ReplaceAll(link, linkbase, "")
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("URL invalide : %v", err)
	}
	// la je retourn la base avec un / a la fin pour faciliter la suite
	base = filepath.Clean(base)

	// Si le chemin est vide ou se termine par "/", traiter comme "index.html"
	if filename == "" || strings.HasSuffix(parsedURL.Path, "/") {
		filename = filepath.Join(filename, "index.html")
	}

	// Gérer l'extension
	// extension := path.Ext(filename)
	// if extension == ""  {
	// 	filename += ".html" // Ajouter une extension par défaut
	// }

	// Construire le chemin complet
	cheminComplet := filepath.Join(base, filename)

	// Créer les répertoires nécessaires
	// err = os.MkdirAll(filepath.Dir(cheminComplet), os.ModePerm)
	// if err != nil {
	// 	return "", fmt.Errorf("Erreur lors de la création des répertoires : %v", err)
	// }

	return cheminComplet, nil
}

func extractLinks(body io.Reader) ([]Link, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("erreur d'analyse HTML : %v", err)
	}

	var links []Link
	var crawler func(*html.Node)

	// Fonction récursive pour parcourir l'arbre HTML
	crawler = func(node *html.Node) {
		// Trouver les nœuds pertinents
		if node.Type == html.ElementNode {
			// Pour les balises ayant href ou src
			if node.Data == "a" || node.Data == "link" || node.Data == "script" || node.Data == "img" {
				for _, attr := range node.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						links = append(links, Link{Href: attr.Val})
					}
				}
			}

			// Gérer les styles inline
			for _, attr := range node.Attr {
				if attr.Key == "style" {
					links = append(links, extractBackgroundImagesFromCSS(attr.Val)...)
				}
			}

			// Gérer les balises <style>
			if node.Data == "style" && node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
				links = append(links, extractBackgroundImagesFromCSS(node.FirstChild.Data)...)
			}
		}

		// Récursion sur les enfants
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(doc)
	return links, nil
}

func extractBackgroundImagesFromCSS(cssContent string) []Link {
	var links []Link
	// Nouvelle expression régulière pour capturer toutes les URLs dans background-image
	re := regexp.MustCompile(`url\(['"]?([^'"\)]+)['"]?\)`)
	matches := re.FindAllStringSubmatch(cssContent, -1)

	// Parcourt toutes les correspondances
	for _, match := range matches {
		if len(match) > 1 {
			links = append(links, Link{Href: strings.TrimSpace(match[1])})
		}
	}
	return links
}


func resolveURL(base, href string) (string, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("URL de base invalide : %v", err)
	}

	resolvedURL, err := baseURL.Parse(href)
	if err != nil {
		return "", fmt.Errorf("URL cible invalide : %v", err)
	}
	if !strings.Contains(resolvedURL.String(), base) && strings.HasPrefix(href, "http") {
		return "", fmt.Errorf("URL cible invalide : %v", href)
	}

	return resolvedURL.String(), nil
}

func Link_log(link Link) {
	// Ouverture du fichier avec les bons drapeaux
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close() // S'assurer que le fichier est fermé

	// Écriture dans le fichier
	_, err = file.WriteString(link.Href + "\n") // Ajouter un saut de ligne
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
	}
}
