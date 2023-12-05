package footlocker

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns struct containing product data
func (r FootlockerRequest) GetProductData(proxy string) (FootlockerData, error) {
	if !r.Valid() {
		return FootlockerData{}, fmt.Errorf("invalid request")
	}

	url := getSearchURL(r.SKU)
	req := utils.NewRequest(url, proxy)
	resp, err := utils.Request(req)
	if err != nil {
		return FootlockerData{}, err
	} else if resp.StatusCode != 200 {
		return FootlockerData{}, fmt.Errorf("error retrieving product data: error code %d", resp.StatusCode)
	}

	json, err := utils.FilterHTML(resp.Body, dataSelector)
	if err != nil {
		return FootlockerData{}, err
	}

	return populateFootlockerData([]byte(json))
}
