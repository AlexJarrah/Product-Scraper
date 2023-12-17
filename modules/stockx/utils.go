package stockx

import (
	"encoding/json"
	"fmt"
	"net/url"

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
	req := utils.NewRequest(u, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return err
	}

	url, err := url.Parse(u)
	if err != nil {
		return err
	}

	cookies := resp.Request.Jar.Cookies(url)
	for _, cookie := range cookies {
		if cookie.Name == "stockx_device_id" {
			deviceID = cookie.Value
			return nil
		}
	}

	return fmt.Errorf("stockx_device_id not found")
}

func unmarshal(data []byte) ([]StockXProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	products := resp.Data.Browse.Results.Edges
	return products, nil
}
