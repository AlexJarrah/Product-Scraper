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
	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return "", err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	return string(resp.Body), nil
}
