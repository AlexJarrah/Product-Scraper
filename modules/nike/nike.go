package nike

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns product data for all Nike products that match the given search query
func FetchNikeProducts(query, proxy string) ([]NikeProduct, error) {
	if query == "" {
		return nil, fmt.Errorf("invalid search query")
	}

	url := getSearchURL(query)

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return nil, err
	}

	json, err := utils.FilterHTML(string(resp.Body), dataSelector)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(json))
}
