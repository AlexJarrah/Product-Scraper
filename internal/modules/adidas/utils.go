package adidas

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Returns the API endpoint to return product data for the provided SKUs
func getAPIEndpoint(skus []string) string {
	return fmt.Sprintf(apiEndpoint, strings.Join(skus, ","))
}

// Parse JSON data into struct
func unmarshal(data []byte) ([]AdidasProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
