package walmart

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AlexJarrah/Product-Scraper/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

func requestProduct(sku, proxy string) (string, error) {
	url := fmt.Sprintf(productEndpoint, sku)

	req := utils.NewRequest(url, proxy)
	resp, err := utils.Request(req)
	if err != nil {
		return "", err
	}

	return resp.Body, err
}

func parseHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	sel := doc.Find(`[id="__NEXT_DATA__"][type="application/json"]`)
	value := sel.Text()

	if value == "" {
		return "", fmt.Errorf("no data found")
	}

	return value, nil
}

func populateCollections(jsonData string) (WalmartProductData, error) {
	jsonData = gjson.Get(jsonData, "props.pageProps.initialData.data.product").String()
	fmt.Println(jsonData)

	data := WalmartProductData{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return WalmartProductData{}, err
	}

	return data, nil
}
