package footlocker

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns product data for the specified product SKU
//
// Note: not currently working; blocked by datadome.
func FetchFootlockerProduct(sku, proxy string) (FootlockerProduct, error) {
	if sku == "" {
		return FootlockerProduct{}, fmt.Errorf("invalid sku")
	}

	url := getAPIEndpoint(sku)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return FootlockerProduct{}, err
	}

	fmt.Println(resp.Body)

	return unmarshal([]byte(resp.Body))
}
