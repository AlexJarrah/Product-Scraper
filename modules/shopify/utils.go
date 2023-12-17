package shopify

import (
	"fmt"
	"net/http"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Formats the provided product URL as the product JSON API endpoint
func getProductAPIEndpoint(url string) string {
	return fmt.Sprintf("%s.json", url)
}

// Retrieves the JSON data of the given URL using the provided proxy
func fetchProductJSON(url, proxy string) (string, error) {
	request := utils.NewRequest(url, proxy)

	response, err := utils.Request(request)
	if err != nil {
		return "", fmt.Errorf("error making HTTP request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	return response.Body, nil
}
