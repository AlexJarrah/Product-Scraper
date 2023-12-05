package snipes

type SnipesRequest struct {
	SKU string `json:"sku"`
}

type SnipesSearchData struct {
	Results []struct {
		Hits []struct {
			GtmImpressionData string `json:"gtmImpressionData"`
			RefinementColor   string `json:"refinementColor"`
			SortDate          int    `json:"sortDate"`
			Categories        [][]struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"categories"`
			VendorStyle          string `json:"vendorStyle"`
			ObjectID             string `json:"objectID"`
			ID                   string `json:"id"`
			SizeBucketRefinement []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"sizeBucketRefinement"`
			Price struct {
				Usd int `json:"USD"`
			} `json:"price"`
			Size []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"size"`
			PrimaryCategoryID string `json:"primary_category_id"`
			InStock           bool   `json:"in_stock"`
			Name              string `json:"name"`
			OnlineFlag        bool   `json:"onlineFlag"`
			Brand             string `json:"brand"`
			PriceData         struct {
				Sales string `json:"sales"`
			} `json:"priceData"`
			AgeGroup string `json:"ageGroup"`
			Variant  bool   `json:"variant"`
			Margin   int    `json:"margin"`
			Gtm      struct {
				ProductOutOfStock bool   `json:"productOutOfStock"`
				Category          string `json:"category"`
				ProductVariant    string `json:"productVariant"`
				SelectVariant     string `json:"selectVariant"`
			} `json:"gtm"`
			Division         string `json:"division"`
			ShortDescription any    `json:"short_description"`
			Master           bool   `json:"master"`
			PrimaryCategory  struct {
				Num0 string `json:"0"`
				Num1 string `json:"1"`
				Num2 string `json:"2"`
			} `json:"__primary_category"`
			Gender string `json:"gender"`
			Badges []struct {
				Name  string `json:"name"`
				Bg    string `json:"bg"`
				Color string `json:"color"`
			} `json:"badges,omitempty"`
			ProductType string `json:"productType"`
			ImageGroups []struct {
				Type   string `json:"_type"`
				Images []struct {
					Type        string `json:"_type"`
					Alt         string `json:"alt"`
					DisBaseLink string `json:"dis_base_link"`
					Title       string `json:"title"`
				} `json:"images"`
				ViewType string `json:"view_type"`
			} `json:"image_groups"`
			LongDescription string   `json:"long_description"`
			Tags            []string `json:"_tags"`
			URL             string   `json:"url"`
			MapPrice        bool     `json:"mapPrice"`
			ColorCount      int      `json:"colorCount"`
		} `json:"hits"`
		NbHits              int    `json:"nbHits"`
		HitsPerPage         int    `json:"hitsPerPage"`
		Index               string `json:"index"`
		Page                int    `json:"page"`
		NbPages             int    `json:"nbPages"`
		ProcessingTimeMS    int    `json:"processingTimeMS"`
		ProcessingTimingsMS struct {
			Request struct {
				RoundTrip int `json:"roundTrip"`
			} `json:"_request"`
		} `json:"processingTimingsMS"`
		ExhaustiveNbHits bool `json:"exhaustiveNbHits"`
		ExhaustiveTypo   bool `json:"exhaustiveTypo"`
		FacetsStats      struct {
			PriceUSD struct {
				Min int     `json:"min"`
				Max int     `json:"max"`
				Avg float64 `json:"avg"`
				Sum int     `json:"sum"`
			} `json:"price.USD"`
		} `json:"facets_stats"`
		Query   string `json:"query"`
		QueryID string `json:"queryID"`
		Params  string `json:"params"`
		Facets  struct {
			Brand struct {
				Jordan int `json:"Jordan"`
				Nike   int `json:"Nike"`
			} `json:"brand"`
			Gender struct {
				MenS   int `json:"Men's"`
				WomenS int `json:"Women's"`
			} `json:"gender"`
			AgeGroup struct {
				Adult int `json:"Adult"`
			} `json:"ageGroup"`
			Division struct {
				Shoes int `json:"Shoes"`
			} `json:"division"`
			PriceUSD struct {
				Num135 int `json:"135"`
				Num140 int `json:"140"`
			} `json:"price.USD"`
			ProductType struct {
				Basketball int `json:"Basketball"`
			} `json:"productType"`
			RefinementColor struct {
				Beige int `json:"Beige"`
				Brown int `json:"Brown"`
			} `json:"refinementColor"`
			SizeBucketRefinementName struct {
				MensShoes10   int `json:"mens-shoes-10"`
				MensShoes105  int `json:"mens-shoes-10.5"`
				MensShoes11   int `json:"mens-shoes-11"`
				MensShoes115  int `json:"mens-shoes-11.5"`
				MensShoes13   int `json:"mens-shoes-13"`
				MensShoes14   int `json:"mens-shoes-14"`
				MensShoes75   int `json:"mens-shoes-7.5"`
				MensShoes8    int `json:"mens-shoes-8"`
				MensShoes85   int `json:"mens-shoes-8.5"`
				MensShoes9    int `json:"mens-shoes-9"`
				MensShoes95   int `json:"mens-shoes-9.5"`
				WomensShoes10 int `json:"womens-shoes-10"`
				WomensShoes11 int `json:"womens-shoes-11"`
				WomensShoes12 int `json:"womens-shoes-12"`
				WomensShoes6  int `json:"womens-shoes-6"`
				WomensShoes65 int `json:"womens-shoes-6.5"`
				WomensShoes7  int `json:"womens-shoes-7"`
				WomensShoes75 int `json:"womens-shoes-7.5"`
				WomensShoes8  int `json:"womens-shoes-8"`
				WomensShoes85 int `json:"womens-shoes-8.5"`
				WomensShoes9  int `json:"womens-shoes-9"`
				WomensShoes95 int `json:"womens-shoes-9.5"`
			} `json:"sizeBucketRefinement.name"`
		} `json:"facets"`
		RenderingContent struct {
		} `json:"renderingContent"`
		Exhaustive struct {
			FacetsCount bool `json:"facetsCount"`
			NbHits      bool `json:"nbHits"`
			Typo        bool `json:"typo"`
		} `json:"exhaustive"`
		ExhaustiveFacetsCount bool `json:"exhaustiveFacetsCount"`
	} `json:"results"`
}
type SnipesData struct {
	Context     string `json:"@context"`
	Type        string `json:"@type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MPN         string `json:"mpn"`
	SKU         string `json:"sku"`
	URL         string `json:"url"`
	Brand       struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"brand"`
	Image  []string `json:"image"`
	Offers struct {
		URL           string `json:"url"`
		Type          string `json:"@type"`
		PriceCurrency string `json:"priceCurrency"`
		Price         string `json:"price"`
		Availability  string `json:"availability"`
		ItemCondition string `json:"itemCondition"`
	} `json:"offers"`
	AdditionalProperty []struct {
		Type       string `json:"@type"`
		PropertyID string `json:"propertyID"`
		Value      string `json:"value"`
	} `json:"additionalProperty"`
	ReleaseDate string `json:"releaseDate"`
	Color       string `json:"color"`
	ID          string `json:"@id"`
}
