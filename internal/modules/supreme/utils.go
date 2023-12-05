package supreme

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func requestCollection() (string, error) {
	resp, err := http.Get(collectionsEndpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), err
}

func parseHTML(html string) (json []byte, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	sel := doc.Find("[class=js-first-all-products-json]")
	value := sel.Text()

	if value == "" {
		return nil, err
	}

	return []byte(value), nil
}

func populateCollections(jsonData []byte) (data SupremeCollectionData, err error) {
	data = SupremeCollectionData{}

	if err = json.Unmarshal(jsonData, &data); err != nil {
		return SupremeCollectionData{}, err
	}

	return data, nil
}

// PopulateStruct takes a struct type and JSON data, and populates an instance of the struct.
func PopulateStruct(structType reflect.Type, jsonData []byte) (interface{}, error) {
	instance := reflect.New(structType).Interface()

	if err := json.Unmarshal(jsonData, instance); err != nil {
		return nil, err
	}

	return instance, nil
}
