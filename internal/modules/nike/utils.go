package nike

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

// Returns search URL for the specified SKU
func getSearchURL(sku string) string {
	return fmt.Sprintf(searchEndpoint, sku, sku)
}

// Parse JSON data into struct
func parseData(jsonData string) (NikeData, error) {
	j := gjson.Get(jsonData, "props.pageProps.initialState.Wall.NikeDatas.0")
	if !j.Exists() {
		return NikeData{}, fmt.Errorf("invalid response")
	}

	data := NikeData{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return NikeData{}, err
	}

	return data, nil
}

func populateNikeData(jsonData string) (NikeData, error) {
	data := NikeData{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return NikeData{}, err
	}

	return data, nil
}
