package adidas

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Checks if request is valid
func (r AdidasRequest) Valid() bool {
	return len(r.SKUs) > 0
}

// Returns API URL for the specified SKUs
func getAPIURL(skus []string) string {
	var list string
	for _, sku := range skus {
		list = fmt.Sprintf("%s,%s", list, sku)
	}
	list = strings.TrimPrefix(list, ",")

	return fmt.Sprintf(apiEndpoint, list)
}

// Parse JSON data into struct
func populateAdidasData(body []byte) (AdidasData, error) {
	data := AdidasData{}
	if err := json.Unmarshal(body, &data); err != nil {
		return AdidasData{}, err
	}

	return data, nil
}
