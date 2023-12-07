package solebox

import (
	"fmt"
	"strings"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
	"github.com/PuerkitoBio/goquery"
)

// Returns product data for all Solebox products that match the given search query
func FetchSoleboxProducts(query, proxy string) ([]SoleboxProduct, error) {
	if query == "" {
		return nil, fmt.Errorf("invalid search query")
	}

	url := getSearchURL(query)
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Body))
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
