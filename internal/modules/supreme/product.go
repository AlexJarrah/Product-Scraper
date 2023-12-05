package supreme

func GetSupremeProduct(proxy string) (SupremeCollectionData, error) {
	html, err := requestCollection()
	if err != nil {
		return SupremeCollectionData{}, err
	}

	json, err := parseHTML(html)
	if err != nil {
		return SupremeCollectionData{}, err
	}

	data, err := populateCollections(json)
	if err != nil {
		return SupremeCollectionData{}, err
	}

	return data, nil
}
