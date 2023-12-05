package solebox

import (
	"encoding/json"
	"fmt"
)

// Checks if request is valid
func (r SoleboxRequest) Valid() bool {
	return r.SKU != ""
}

// Returns product URL for the specified SKU
func getProductURL(sku string) string {
	return fmt.Sprintf(apiEndpoint, sku)
}

// Parse JSON data into struct
func populateSoleboxData(jsonData []byte) (SoleboxData, error) {
	data := SoleboxData{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return SoleboxData{}, err
	}

	return data, nil
}
