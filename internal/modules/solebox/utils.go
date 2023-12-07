package solebox

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Returns search URL for the specified query
func getSearchURL(query string) string {
	q := url.QueryEscape(query)
	return fmt.Sprintf(searchEndpoint, q)
}

func unmarshal(data []byte) (SoleboxProduct, error) {
	resp := SoleboxProduct{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return SoleboxProduct{}, err
	}

	return resp, nil
}
