package supreme

import (
	"encoding/json"
	"fmt"
)

// Returns API endpoint to return all product datea for the season
func getSeasonAPIEndpoint(season string) string {
	return fmt.Sprintf(seasonAPIEndpoint, season)
}

// Returns endpoint to return all current shop data
func getAllCollectionsEndpoint() string {
	return collectionsEndpoint
}

func unmarshalCollections(data []byte) (SupremeCollections, error) {
	resp := responseCollections{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return SupremeCollections{}, err
	}

	return SupremeCollections(resp), nil
}

func unmarshalSeason(data []byte) (SupremeSeason, error) {
	resp := responseSeason{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return SupremeSeason{}, err
	}

	return resp.Result.Data, nil
}
