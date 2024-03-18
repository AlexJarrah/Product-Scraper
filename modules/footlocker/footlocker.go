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

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return FootlockerProduct{}, err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return FootlockerProduct{}, err
	}

	return unmarshal([]byte(resp.Body))
}
