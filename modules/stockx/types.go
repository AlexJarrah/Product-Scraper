package stockx

import "time"

type StockXProduct struct {
	ObjectID string `json:"objectId"`
	Node     struct {
		ID              string   `json:"id"`
		Name            string   `json:"name"`
		URLKey          string   `json:"urlKey"`
		Title           string   `json:"title"`
		Brand           string   `json:"brand"`
		Description     string   `json:"description"`
		Model           string   `json:"model"`
		Condition       string   `json:"condition"`
		ProductCategory string   `json:"productCategory"`
		ListingType     string   `json:"listingType"`
		Gender          string   `json:"gender"`
		BrowseVerticals []string `json:"browseVerticals"`
		Media           struct {
			ThumbURL      string `json:"thumbUrl"`
			SmallImageURL string `json:"smallImageUrl"`
		} `json:"media"`
		Favorite      any `json:"favorite"`
		ProductTraits []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"productTraits"`
		Variants []struct {
			ID       string `json:"id"`
			Hidden   bool   `json:"hidden"`
			Favorite any    `json:"favorite"`
			Traits   struct {
				Size string `json:"size"`
			} `json:"traits"`
		} `json:"variants"`
		Market struct {
			CurrencyCode string `json:"currencyCode"`
			BidAskData   struct {
				LowestAsk          int       `json:"lowestAsk"`
				HighestBid         int       `json:"highestBid"`
				LastHighestBidTime time.Time `json:"lastHighestBidTime"`
				LastLowestAskTime  time.Time `json:"lastLowestAskTime"`
			} `json:"bidAskData"`
			State struct {
				NumberOfCustodialAsks int `json:"numberOfCustodialAsks"`
			} `json:"state"`
			SalesInformation struct {
				LastSale         int       `json:"lastSale"`
				LastSaleDate     time.Time `json:"lastSaleDate"`
				SalesThisPeriod  int       `json:"salesThisPeriod"`
				SalesLastPeriod  int       `json:"salesLastPeriod"`
				ChangeValue      int       `json:"changeValue"`
				ChangePercentage float64   `json:"changePercentage"`
				Volatility       float64   `json:"volatility"`
				PricePremium     float64   `json:"pricePremium"`
			} `json:"salesInformation"`
			DeadStock struct {
				Sold         int `json:"sold"`
				AveragePrice int `json:"averagePrice"`
			} `json:"deadStock"`
			Statistics struct {
				Last90Days struct {
					AveragePrice int `json:"averagePrice"`
				} `json:"last90Days"`
			} `json:"statistics"`
		} `json:"market"`
	} `json:"node"`
}

type response struct {
	Data struct {
		Browse struct {
			Results struct {
				Edges []StockXProduct `json:"edges"`
			} `json:"results"`
		} `json:"browse"`
	} `json:"data"`
}
