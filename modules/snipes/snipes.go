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
	req := utils.NewRequest(url, proxy)
	req.Method = "POST"
	req.Body = payload

	resp, err := utils.Request(req)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(resp.Body))
}
