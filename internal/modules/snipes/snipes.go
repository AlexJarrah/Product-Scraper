package snipes

import (
	"errors"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns struct containing product data
func (r SnipesRequest) GetProductData(proxy string) (SnipesSearchData, error) {
	if !r.Valid() {
		return SnipesSearchData{}, errors.New("invalid request")
	}

	url := getAPIURL()
	payload := getAPIPayload(r.SKU)
	req := utils.NewRequest(url, proxy)
	req.Method = "POST"
	req.Body = payload

	resp, err := utils.Request(req)
	if err != nil {
		return SnipesSearchData{}, err
	}

	return populateSnipesData([]byte(resp.Body))
}
