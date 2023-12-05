package adidas

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns struct containing product data
func (r AdidasRequest) GetProductsData(proxy string) (AdidasData, error) {
	if !r.Valid() {
		return AdidasData{}, fmt.Errorf("invalid request")
	}

	url := getAPIURL(r.SKUs)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return AdidasData{}, err
	}

	return populateAdidasData([]byte(resp.Body))
}
