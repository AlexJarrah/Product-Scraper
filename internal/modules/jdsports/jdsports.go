package jdsports

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns struct containing product data
func (r JDSportsRequest) GetProductData(proxy string) (JDSportsData, error) {
	if !r.Valid() {
		return JDSportsData{}, fmt.Errorf("invalid request")
	}

	url := getProductURL(r.SKU)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return JDSportsData{}, err
	}

	return populateJDSportsData([]byte(resp.Body))
}
