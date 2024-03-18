package stockx

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns product market data for all StockX products matching the given search query
func FetchStockXProducts(query, proxy string) ([]StockXProduct, error) {
	if query == "" {
		return nil, fmt.Errorf("invalid search query")
	}

	// Sets deviceID to value of cookie [x-stockx-device-id] if not already set.
	if err := getDeviceID(proxy); err != nil {
		return nil, err
	}

	session, err := utils.NewSession(apiEndpoint, proxy)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	session.OrderedHeaders.Set("apollographql-client-name", "Iron")
	session.OrderedHeaders.Set("apollographql-client-version", clientVersion)
	session.OrderedHeaders.Set("app-platform", "Iron")
	session.OrderedHeaders.Set("app-version", clientVersion)
	session.OrderedHeaders.Set("accept-language", "en-US")
	session.OrderedHeaders.Set("x-stockx-device-id", deviceID)

	body := getRequestBody(query)
	resp, err := session.Post(apiEndpoint, body)
	if err != nil {
		return nil, err
	}

	return unmarshal(resp.Body)
}
