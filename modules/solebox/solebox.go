package solebox

import (
	"bytes"
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/utils"
	"github.com/PuerkitoBio/goquery"
)

// Returns product data for all Solebox products that match the given search query
func FetchSoleboxProducts(query, proxy string) ([]SoleboxProduct, error) {
	if query == "" {
		return nil, fmt.Errorf("invalid search query")
	}

	url := getSearchURL(query)

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body))
	if err != nil {
		return nil, err
	}

	var res []SoleboxProduct

	elements := doc.Find(dataSelector)
	for _, el := range elements.Nodes {
		for _, attr := range el.Attr {
			if attr.Key == dataAttribute {
				if v, err := unmarshal([]byte(attr.Val)); err == nil {
					res = append(res, v)
				}
			}
		}
	}

	return res, nil
}
