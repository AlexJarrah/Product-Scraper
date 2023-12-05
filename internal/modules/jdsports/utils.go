package jdsports

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Checks if request is valid
func (r JDSportsRequest) Valid() bool {
	return r.SKU != ""
}

// Returns product URL for the specified SKU
func getProductURL(sku string) string {
	return fmt.Sprintf(apiEndpoint, strings.ReplaceAll(sku, "-", "_"))
}

// Parse JSON data into struct
func populateJDSportsData(body []byte) (JDSportsData, error) {
	data := JDSportsData{}
	if err := json.Unmarshal(body, &data); err != nil {
		return JDSportsData{}, err
	}

	return data, nil
}
