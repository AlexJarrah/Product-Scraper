package snipes

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns struct containing product data
func FetchSnipesProducts(query, proxy string) ([]SnipesProduct, error) {
	if query == "" {
		return nil, fmt.Errorf("invalid search query")
	}

	url := getAPIURL()
	payload := getAPIPayload(query)

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	resp, err := session.Post(url, payload)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(resp.Body))
}
