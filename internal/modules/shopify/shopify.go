package shopify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

func Shopify(url, proxy string) (ShopifyProduct, error) {
	if !utils.IsValidURL(url) {
		return ShopifyProduct{}, fmt.Errorf("invalid url")
	}

	url, err := utils.RemoveURLParameters(url)
	if err != nil {
		return ShopifyProduct{}, err
	}

	apiEndpoint := shopifyGetAPIEndpoint(url)
	data, err := shopifyGetProductJSON(apiEndpoint)
	if err != nil {
		return ShopifyProduct{}, err
	}

	var res ShopifyProduct
	if err := json.Unmarshal(data, &res); err != nil {
		return ShopifyProduct{}, err
	}

	return res, nil
}

func shopifyGetAPIEndpoint(url string) string {
	return fmt.Sprintf("%s.json", url)
}

func shopifyGetProductJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
