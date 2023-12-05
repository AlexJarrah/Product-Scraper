package footlocker

type FootlockerRequest struct {
	SKU string `json:"sku"`
}

type FootlockerData struct {
	Context         string `json:"@context"`
	Type            string `json:"@type"`
	ID              string `json:"@id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	Brand           string `json:"brand"`
	Model           string `json:"model"`
	SKU             string `json:"sku"`
	URL             string `json:"url"`
	ItemCondition   string `json:"itemCondition"`
	AggregateRating struct {
		Type        string `json:"@type"`
		RatingValue string `json:"ratingValue"`
		ReviewCount int    `json:"reviewCount"`
	} `json:"aggregateRating"`
	Review []struct {
		Type          string `json:"@type"`
		Author        string `json:"author"`
		DatePublished string `json:"datePublished"`
		ReviewBody    string `json:"reviewBody"`
		Name          string `json:"name"`
		ReviewRating  struct {
			Type        string `json:"@type"`
			BestRating  int    `json:"bestRating"`
			RatingValue int    `json:"ratingValue"`
			WorstRating int    `json:"worstRating"`
		} `json:"reviewRating"`
	} `json:"review"`
	Offers []struct {
		Type         string `json:"@type"`
		SKU          string `json:"sku"`
		Price        int    `json:"price"`
		Seller       string `json:"seller"`
		Availability string `json:"availability"`
	} `json:"offers"`
}
