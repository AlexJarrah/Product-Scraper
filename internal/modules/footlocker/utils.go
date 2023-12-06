package footlocker

import (
	"encoding/json"
	"fmt"
)

// Returns product URL for the specified SKU
func getProductURL(sku string) string {
	return fmt.Sprintf(productEndpoint, sku)
}

func unmarshal(data []byte) (FootlockerProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return FootlockerProduct{}, err
	}

	return FootlockerProduct(resp), nil
}
