package adidas

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns product data from Adidas for the specified SKUs
func FetchAdidasProducts(skus []string, proxy string) ([]AdidasProduct, error) {
	if len(skus) == 0 {
		return nil, fmt.Errorf("invalid sku list")
	}

	url := getAPIEndpoint(skus)

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(resp.Body))
}
