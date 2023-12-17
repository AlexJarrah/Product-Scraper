package nike

type NikeProduct struct {
	CardType         string `json:"cardType"`
	CloudProductID   string `json:"cloudProductId"`
	ColorDescription string `json:"colorDescription"`
	Colorways        []struct {
		ColorDescription string `json:"colorDescription"`
		Images           struct {
			PortraitURL string `json:"portraitURL"`
			SquarishURL string `json:"squarishURL"`
		} `json:"images"`
		PdpURL string `json:"pdpUrl"`
		Price  struct {
			Currency               string  `json:"currency"`
			CurrentPrice           float32 `json:"currentPrice"`
			Discounted             bool    `json:"discounted"`
			EmployeePrice          float32 `json:"employeePrice"`
			FullPrice              float32 `json:"fullPrice"`
			MinimumAdvertisedPrice any     `json:"minimumAdvertisedPrice"`
		} `json:"price"`
		CloudProductID    string `json:"cloudProductId"`
		InStock           bool   `json:"inStock"`
		IsComingSoon      bool   `json:"isComingSoon"`
		IsBestSeller      bool   `json:"isBestSeller"`
		IsExcluded        bool   `json:"isExcluded"`
		IsJPStrikethrough any    `json:"isJPStrikethrough"`
		IsLaunch          bool   `json:"isLaunch"`
		IsMemberExclusive bool   `json:"isMemberExclusive"`
		IsNew             bool   `json:"isNew"`
		Label             string `json:"label"`
		Pid               string `json:"pid"`
		PrebuildID        any    `json:"prebuildId"`
		ProductInstanceID any    `json:"productInstanceId"`
	} `json:"colorways"`
	Customizable      bool   `json:"customizable"`
	HasExtendedSizing bool   `json:"hasExtendedSizing"`
	ID                string `json:"id"`
	Images            struct {
		PortraitURL string `json:"portraitURL"`
		SquarishURL string `json:"squarishURL"`
	} `json:"images"`
	InStock           bool   `json:"inStock"`
	IsComingSoon      bool   `json:"isComingSoon"`
	IsBestSeller      bool   `json:"isBestSeller"`
	IsExcluded        bool   `json:"isExcluded"`
	IsGiftCard        bool   `json:"isGiftCard"`
	IsJersey          bool   `json:"isJersey"`
	IsJPStrikethrough any    `json:"isJPStrikethrough"`
	IsLaunch          bool   `json:"isLaunch"`
	IsMemberExclusive bool   `json:"isMemberExclusive"`
	IsNBA             bool   `json:"isNBA"`
	IsNFL             bool   `json:"isNFL"`
	IsSustainable     bool   `json:"isSustainable"`
	Label             string `json:"label"`
	NbyColorway       any    `json:"nbyColorway"`
	Pid               string `json:"pid"`
	PrebuildID        any    `json:"prebuildId"`
	Price             struct {
		Currency               string  `json:"currency"`
		CurrentPrice           float32 `json:"currentPrice"`
		Discounted             bool    `json:"discounted"`
		EmployeePrice          float32 `json:"employeePrice"`
		FullPrice              float32 `json:"fullPrice"`
		MinimumAdvertisedPrice any     `json:"minimumAdvertisedPrice"`
	} `json:"price"`
	ProductInstanceID any      `json:"productInstanceId"`
	ProductType       string   `json:"productType"`
	Properties        any      `json:"properties"`
	SalesChannel      []string `json:"salesChannel"`
	Subtitle          string   `json:"subtitle"`
	Title             string   `json:"title"`
	URL               string   `json:"url"`
}

type response struct {
	Props struct {
		PageProps struct {
			InitialState struct {
				Wall struct {
					Products []NikeProduct `json:"products"`
				} `json:"Wall"`
			} `json:"initialState"`
		} `json:"pageProps"`
	} `json:"props"`
}
