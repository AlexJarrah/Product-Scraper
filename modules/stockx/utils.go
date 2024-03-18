package stockx

import (
	"encoding/json"
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns request body for the specified query
func getRequestBody(query string) string {
	return fmt.Sprintf(requestBody, query, query)
}

func getDeviceID(proxy string) error {
	if deviceID != "" {
		return nil
	}

	u := homepage

	session, err := utils.NewSession(u, proxy)
	if err != nil {
		return err
	}
	defer session.Close()

	resp, err := session.Get(u)
	if err != nil {
		return err
	}

	for _, cookie := range resp.HttpResponse.Cookies() {
		if cookie.Name == "stockx_device_id" {
			deviceID = cookie.Value
			return nil
		}
	}

	return fmt.Errorf("cookie stockx_device_id not found")
}

func unmarshal(data []byte) ([]StockXProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	products := resp.Data.Browse.Results.Edges
	return products, nil
}
