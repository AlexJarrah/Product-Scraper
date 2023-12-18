package goat

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns product market data for all Goat products matching the given search query
func FetchGoatProducts(query, proxy string) ([]GoatProduct, error) {
	if query == "" {
		return nil, fmt.Errorf("invalid search query")
	}

	url := getSearchURL(query)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(resp.Body))
}
