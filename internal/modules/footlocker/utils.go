package footlocker

import (
	"encoding/json"
	"fmt"
)

// Checks if request is valid
func (r FootlockerRequest) Valid() bool {
	return r.SKU != ""
}

// Returns product URL for the specified SKU
func getSearchURL(sku string) string {
	return fmt.Sprintf(productEndpoint, sku)
}

// Parse JSON data into struct
func populateFootlockerData(body []byte) (FootlockerData, error) {
	data := FootlockerData{}
	if err := json.Unmarshal(body, &data); err != nil {
		return FootlockerData{}, err
	}

	return data, nil
}
