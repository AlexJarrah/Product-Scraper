package nike

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns struct containing NikeData data
func NikeProduct(sku, proxy string) (NikeData, error) {
	if sku == "" {
		return NikeData{}, fmt.Errorf("invalid request")
	}

	url := getSearchURL(sku)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return NikeData{}, err
	}

	json, err := utils.FilterHTML(resp.Body, dataSelector)
	if err != nil {
		return NikeData{}, err
	}

	fmt.Println(json)

	return populateNikeData(json)
}
