package shopify

import (
	"encoding/json"
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Retrieves Shopify product data from the provided URL
func Shopify(url, proxy string) (ShopifyProduct, error) {
	// Validate the URL
	if !utils.IsValidURL(url) {
		return ShopifyProduct{}, fmt.Errorf("invalid url")
	}

	// Remove query parameters from the URL
	url, err := utils.RemoveURLParameters(url)
	if err != nil {
		return ShopifyProduct{}, fmt.Errorf("error parsing url: %w", err)
	}

	// Get the Shopify API endpoint URL
	apiEndpoint := getProductAPIEndpoint(url)

	// Retrieve product JSON data
	data, err := fetchProductJSON(apiEndpoint, proxy)
	if err != nil {
		return ShopifyProduct{}, fmt.Errorf("error getting product data: %w", err)
	}

	// Parse product JSON data
	var res ShopifyProduct
	if err := json.Unmarshal([]byte(data), &res); err != nil {
		return ShopifyProduct{}, fmt.Errorf("error parsing product data: %w", err)
	}

	return res, nil
}
