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

	req := utils.NewRequest(apiEndpoint, proxy)
	req.Method = "POST"
	req.Body = getRequestBody(query)
	req.Headers["apollographql-client-name"] = "Iron"
	req.Headers["apollographql-client-version"] = clientVersion
	req.Headers["app-platform"] = "Iron"
	req.Headers["app-version"] = clientVersion
	req.Headers["accept-language"] = "en-US"
	req.Headers["x-stockx-device-id"] = deviceID

	resp, err := utils.Request(req)
	if err != nil {
		return nil, err
	}

	return unmarshal([]byte(resp.Body))
}
