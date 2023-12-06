package adidas

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns product data from Adidas for the specified SKUs
func FetchAdidasProducts(skus []string, proxy string) ([]AdidasProduct, error) {
	if len(skus) == 0 {
		return nil, fmt.Errorf("invalid sku list")
	}

	url := getAPIEndpoint(skus)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(resp.Body))
}
