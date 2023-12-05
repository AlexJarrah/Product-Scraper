package nike

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns product data for all Nike products that match the given search query
func FetchNikeProducts(query, proxy string) ([]NikeProduct, error) {
	if query == "" {
		return nil, fmt.Errorf("invalid search query")
	}

	url := getSearchURL(query)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return nil, err
	}

	json, err := utils.FilterHTML(resp.Body, dataSelector)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(json))
}
