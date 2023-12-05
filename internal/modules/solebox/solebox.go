package solebox

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns struct containing product data
func (r SoleboxRequest) GetProductData(proxy string) (SoleboxData, error) {
	if !r.Valid() {
		return SoleboxData{}, fmt.Errorf("invalid request")
	}

	url := getProductURL(r.SKU)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return SoleboxData{}, err
	}

	json, err := utils.FilterHTML(resp.Body, dataSelector, dataAttribute)
	if err != nil {
		return SoleboxData{}, err
	}

	return populateSoleboxData([]byte(json))
}
