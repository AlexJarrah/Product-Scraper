package supreme

import (
	"github.com/AlexJarrah/Product-Scraper/utils"
)

// Returns product data for all products in the Supreme collection
func FetchSupremeCollections(proxy string) (SupremeCollections, error) {
	url := getAllCollectionsEndpoint()

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return SupremeCollections{}, err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return SupremeCollections{}, err
	}

	json, err := utils.FilterHTML(string(resp.Body), collectionsJSONSelector)
	if err != nil {
		return SupremeCollections{}, err
	}

	return unmarshalCollections([]byte(json))
}

// Returns product data for every product in the Supreme season lookbook & preview
func FetchSupremeSeason(season, proxy string) (SupremeSeason, error) {
	url := getSeasonAPIEndpoint(season)

	session, err := utils.NewSession(url, proxy)
	if err != nil {
		return SupremeSeason{}, err
	}
	defer session.Close()

	resp, err := session.Get(url)
	if err != nil {
		return SupremeSeason{}, err
	}

	return unmarshalSeason([]byte(resp.Body))
}
