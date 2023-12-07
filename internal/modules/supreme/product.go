package supreme

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/utils"
)

// Returns product data for all products in the Supreme collection
func FetchSupremeCollections(proxy string) (SupremeCollections, error) {
	url := getAllCollectionsEndpoint()
	req := utils.NewRequest(url, proxy)

	resp, err := utils.Request(req)
	if err != nil {
		return SupremeCollections{}, err
	}

	json, err := utils.FilterHTML(resp.Body, collectionsJSONSelector)
	if err != nil {
		return SupremeCollections{}, err
	}

	return unmarshalCollections([]byte(json))
}

// Returns product data for every product in the Supreme season lookbook & preview
func FetchSupremeSeason(season, proxy string) (SupremeSeason, error) {
	url := getSeasonAPIEndpoint(season)
	req := utils.NewRequest(url, proxy)

	fmt.Println(url)

	resp, err := utils.Request(req)
	if err != nil {
		return SupremeSeason{}, err
	}

	return unmarshalSeason([]byte(resp.Body))
}
