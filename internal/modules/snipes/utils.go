package snipes

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Returns API endpoint for the specified query
func getAPIURL() string {
	return apiEndpoint
}

// Returns API payload for the specified query
func getAPIPayload(query string) string {
	q := url.QueryEscape(query)
	return strings.ReplaceAll(apiPayload, "{query}", q)
}

// Parse JSON data into struct
func unmarshal(data []byte) ([]SnipesProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	products := resp.Results[0].Hits
	return products, nil
}
