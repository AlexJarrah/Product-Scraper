package products

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/quo0001/Product-Scraper/utils"
)

func Shopify(url, proxy string) (ShopifyProduct, error) {
	if !utils.IsValidURL(url) {
		return ShopifyProduct{}, fmt.Errorf("invalid url")
	}

	url, err := utils.RemoveURLParameters(url)
	if err != nil {
		return ShopifyProduct{}, err
	}

	apiEndpoint := getAPIEndpoint(url)
	data, err := getProductJSON(apiEndpoint)
	if err != nil {
		return ShopifyProduct{}, err
	}

	var res ShopifyProduct
	if err := json.Unmarshal(data, &res); err != nil {
		return ShopifyProduct{}, err
	}

	return res, nil
}

func getAPIEndpoint(url string) string {
	return fmt.Sprintf("%s.json", url)
}

func getProductJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
