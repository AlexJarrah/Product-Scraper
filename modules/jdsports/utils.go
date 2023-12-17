package jdsports

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Returns product API endpoint for the specified SKU
func getProductAPIEndpoint(sku string) string {
	return fmt.Sprintf(productAPIEndpoint, strings.ReplaceAll(sku, "-", "_"))
}

func unmarshal(data []byte) (JDSportsProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return JDSportsProduct{}, err
	}

	if len(resp.Outfits) == 0 {
		return JDSportsProduct{}, fmt.Errorf("no products found")
	} else if len(resp.Outfits[0].Items) == 0 {
		return JDSportsProduct{}, fmt.Errorf("no products found")
	}

	product := resp.Outfits[0].Items[0]
	return product, nil
}
