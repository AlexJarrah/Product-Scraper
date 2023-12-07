package footlocker

import (
	"encoding/json"
	"fmt"
)

// Returns the API endpoint to provide product data on the specified SKU
func getAPIEndpoint(sku string) string {
	return fmt.Sprintf(productEndpoint, sku)
}

func unmarshal(data []byte) (FootlockerProduct, error) {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return FootlockerProduct{}, err
	}

	return FootlockerProduct(resp), nil
}
