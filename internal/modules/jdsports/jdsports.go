package jdsports

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns product data for the JD Sports product matching the given SKU
func FetchJDSportsProduct(sku, proxy string) (JDSportsProduct, error) {
	if sku == "" {
		return JDSportsProduct{}, fmt.Errorf("invalid sku")
	}

	url := getProductAPIEndpoint(sku)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return JDSportsProduct{}, err
	}

	return unmarshal([]byte(resp.Body))
}
