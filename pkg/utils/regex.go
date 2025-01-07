package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func CheckLink(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération de l'URL: %v", err)
	}
	//defer response.Body.Close()
	return response, nil
}

func CheckURL(url string) bool {
	reg := regexp.MustCompile(`^(https?):\/\/[^\s/$.?#].[^\s]*$`)
	return strings.TrimSpace(url) != "" && reg.MatchString(url)
}
