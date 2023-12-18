package goat

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Returns search URL for the specified SKU
func getSearchURL(query string) string {
	q := url.QueryEscape(query)
	return fmt.Sprintf(searchEndpoint, q)
}

func unmarshal(data []byte) ([]GoatProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	products := resp.Response.Results
	return products, nil
}
