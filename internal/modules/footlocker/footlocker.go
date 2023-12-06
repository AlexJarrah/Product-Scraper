package footlocker

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns product data for the specified product SKU
func FetchFootlockerProduct(sku, proxy string) (FootlockerProduct, error) {
	if sku == "" {
		return FootlockerProduct{}, fmt.Errorf("invalid sku")
	}

	url := getProductURL(sku)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return FootlockerProduct{}, err
	}

	json, err := utils.FilterHTML(resp.Body, dataSelector)
	if err != nil {
		return FootlockerProduct{}, err
	}

	return unmarshal([]byte(json))
}
