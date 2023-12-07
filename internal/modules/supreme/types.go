package supreme

import "time"

type SupremeCollections struct {
	AllProductsCount int `json:"allProductsCount"`
	Products         []struct {
		ID          int64    `json:"id"`
		Title       string   `json:"title"`
		Handle      string   `json:"handle"`
		URL         string   `json:"url"`
		Tags        []string `json:"tags"`
		ProductType string   `json:"product_type"`
		Image       string   `json:"image"`
		Images      []struct {
			Src    string `json:"src"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			ID     int64  `json:"id"`
			Alt    string `json:"alt"`
		} `json:"images"`
		Price       int      `json:"price"`
		Description string   `json:"description"`
		Type        string   `json:"type"`
		Available   bool     `json:"available"`
		Collection  string   `json:"collection"`
		Options     []string `json:"options"`
		Variants    []struct {
			ID                     int64    `json:"id"`
			Title                  string   `json:"title"`
			Option1                string   `json:"option1"`
			Option2                any      `json:"option2"`
			Option3                any      `json:"option3"`
			Sku                    string   `json:"sku"`
			RequiresShipping       bool     `json:"requires_shipping"`
			Taxable                bool     `json:"taxable"`
			FeaturedImage          any      `json:"featured_image"`
			Available              bool     `json:"available"`
			Name                   string   `json:"name"`
			PublicTitle            string   `json:"public_title"`
			Options                []string `json:"options"`
			Price                  int      `json:"price"`
			Weight                 int      `json:"weight"`
			CompareAtPrice         any      `json:"compare_at_price"`
			InventoryManagement    string   `json:"inventory_management"`
			Barcode                string   `json:"barcode"`
			RequiresSellingPlan    bool     `json:"requires_selling_plan"`
			SellingPlanAllocations []any    `json:"selling_plan_allocations"`
			QuantityRule           struct {
				Min       int `json:"min"`
				Max       any `json:"max"`
				Increment int `json:"increment"`
			} `json:"quantity_rule"`
		} `json:"variants"`
		SelectedVariant struct {
			ID                     int64    `json:"id"`
			Title                  string   `json:"title"`
			Option1                string   `json:"option1"`
			Option2                any      `json:"option2"`
			Option3                any      `json:"option3"`
			Sku                    string   `json:"sku"`
			RequiresShipping       bool     `json:"requires_shipping"`
			Taxable                bool     `json:"taxable"`
			FeaturedImage          any      `json:"featured_image"`
			Available              bool     `json:"available"`
			Name                   string   `json:"name"`
			PublicTitle            string   `json:"public_title"`
			Options                []string `json:"options"`
			Price                  int      `json:"price"`
			Weight                 int      `json:"weight"`
			CompareAtPrice         any      `json:"compare_at_price"`
			InventoryManagement    string   `json:"inventory_management"`
			Barcode                string   `json:"barcode"`
			RequiresSellingPlan    bool     `json:"requires_selling_plan"`
			SellingPlanAllocations []any    `json:"selling_plan_allocations"`
			QuantityRule           struct {
				Min       int `json:"min"`
				Max       any `json:"max"`
				Increment int `json:"increment"`
			} `json:"quantity_rule"`
		} `json:"selectedVariant"`
	} `json:"products"`
}

type SupremeSeason struct {
	Lookbooks struct {
		Nodes []struct {
			Season struct {
				Title string `json:"title"`
				Slug  struct {
					Current string `json:"current"`
				} `json:"slug"`
			} `json:"season"`
			PublishedAt time.Time `json:"publishedAt"`
			Slug        struct {
				Current string `json:"current"`
			} `json:"slug"`
		} `json:"nodes"`
	} `json:"lookbooks"`
	Preview struct {
		ID          string    `json:"id"`
		PublishedAt time.Time `json:"publishedAt"`
		Season      struct {
			Title string `json:"title"`
			Slug  struct {
				Current string `json:"current"`
			} `json:"slug"`
		} `json:"season"`
		Groups []struct {
			Title     string `json:"title"`
			Thumbnail struct {
				Asset struct {
					ID              string `json:"_id"`
					GatsbyImageData struct {
						Images struct {
							Sources  []any `json:"sources"`
							Fallback struct {
								Src    string `json:"src"`
								SrcSet string `json:"srcSet"`
								Sizes  string `json:"sizes"`
							} `json:"fallback"`
						} `json:"images"`
						Layout          string `json:"layout"`
						BackgroundColor string `json:"backgroundColor"`
						Width           int    `json:"width"`
						Height          int    `json:"height"`
					} `json:"gatsbyImageData"`
				} `json:"asset"`
				Hotspot  any    `json:"hotspot"`
				Crop     any    `json:"crop"`
				Typename string `json:"__typename"`
				Type     string `json:"_type"`
				Key      any    `json:"_key"`
			} `json:"thumbnail"`
			Products []struct {
				Title    string `json:"title"`
				Category struct {
					Title string `json:"title"`
					Slug  struct {
						Current string `json:"current"`
					} `json:"slug"`
				} `json:"category"`
				Slug struct {
					Current string `json:"current"`
				} `json:"slug"`
				Variants []struct {
					Title                 string `json:"title"`
					HideFromGroupCarousel any    `json:"hideFromGroupCarousel"`
					FilteredImages        []struct {
						Asset struct {
							ID              string `json:"_id"`
							GatsbyImageData struct {
								Images struct {
									Sources  []any `json:"sources"`
									Fallback struct {
										Src    string `json:"src"`
										SrcSet string `json:"srcSet"`
										Sizes  string `json:"sizes"`
									} `json:"fallback"`
								} `json:"images"`
								Layout string  `json:"layout"`
								Width  float32 `json:"width"`
								Height float32 `json:"height"`
							} `json:"gatsbyImageData"`
						} `json:"asset"`
					} `json:"filteredImages"`
				} `json:"variants"`
				FilteredImages []struct {
					HideFromGroupCarousel bool `json:"hideFromGroupCarousel"`
					Alt                   any  `json:"alt"`
					Asset                 struct {
						ID              string `json:"_id"`
						GatsbyImageData struct {
							Images struct {
								Sources  []any `json:"sources"`
								Fallback struct {
									Src    string `json:"src"`
									SrcSet string `json:"srcSet"`
									Sizes  string `json:"sizes"`
								} `json:"fallback"`
							} `json:"images"`
							Layout string  `json:"layout"`
							Width  float32 `json:"width"`
							Height float32 `json:"height"`
						} `json:"gatsbyImageData"`
					} `json:"asset"`
				} `json:"filteredImages"`
			} `json:"products"`
		} `json:"groups"`
	} `json:"preview"`
}

type responseCollections SupremeCollections

type responseSeason struct {
	Result struct {
		Data SupremeSeason `json:"data"`
	} `json:"result"`
}
