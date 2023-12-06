package adidas

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Returns the API endpoint to provide product data on the specified SKUs
func getAPIEndpoint(skus []string) string {
	return fmt.Sprintf(apiEndpoint, strings.Join(skus, ","))
}

func unmarshal(data []byte) ([]AdidasProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
