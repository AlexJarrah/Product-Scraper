package walmart

func GetWalmartProduct(r WalmartProductRequest, proxy string) (WalmartProductData, error) {
	html, err := requestProduct(r.SKU)
	if err != nil {
		return WalmartProductData{}, err
	}

	json, err := parseHTML(html)
	if err != nil {
		return WalmartProductData{}, err
	}

	data, err := populateCollections(json)
	if err != nil {
		return WalmartProductData{}, err
	}

	return data, nil
}
