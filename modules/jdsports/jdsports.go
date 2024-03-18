package jdsports

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns product data for the JD Sports product matching the given SKU
func FetchJDSportsProduct(sku, proxy string) (JDSportsProduct, error) {
	if sku == "" {
		return JDSportsProduct{}, fmt.Errorf("invalid sku")
	}

	url := getProductAPIEndpoint(sku)

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return JDSportsProduct{}, err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return JDSportsProduct{}, err
	}

	return unmarshal([]byte(resp.Body))
}
