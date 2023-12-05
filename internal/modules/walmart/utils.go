package walmart

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func requestProduct(sku string) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf(productEndpoint, sku)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
	req.Header.Add("host", "www.walmart.com")

	resp, err := client.Do(req)
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

	sel := doc.Find(`[id="__NEXT_DATA__"][type="application/json"]`)
	value := sel.Text()

	if value == "" {
		return nil, fmt.Errorf("no data found")
	}

	return []byte(value), nil
}

func populateCollections(jsonData []byte) (data WalmartProductData, err error) {
	data = WalmartProductData{}

	if err = json.Unmarshal(jsonData, &data); err != nil {
		return WalmartProductData{}, err
	}

	return data, nil
}
