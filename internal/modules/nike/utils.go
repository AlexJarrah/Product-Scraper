package nike

import (
	"encoding/json"
	"fmt"
)

// Returns search URL for the specified SKU
func getSearchURL(query string) string {
	return fmt.Sprintf(searchEndpoint, query, query)
}

func unmarshal(data []byte) ([]NikeProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	products := resp.Props.PageProps.InitialState.Wall.Products
	return products, nil
}
