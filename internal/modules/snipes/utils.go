package snipes

import (
	"encoding/json"
	"strings"
)

// Checks if request is valid
func (r SnipesRequest) Valid() bool {
	return r.SKU != ""
}

// Returns API endpoint for the specified SKU
func getAPIURL() string {
	return apiEndpoint
}

// Returns API payload for the specified SKU
func getAPIPayload(sku string) string {
	return strings.ReplaceAll(apiPayload, "{sku}", sku)
}

// Parse JSON data into struct
func populateSnipesData(jsonData []byte) (SnipesSearchData, error) {
	data := SnipesSearchData{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return SnipesSearchData{}, err
	}

	return data, nil
}
