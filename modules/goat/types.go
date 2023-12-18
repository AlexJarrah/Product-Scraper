package goat

type GoatProduct struct {
	MatchedTerms []string `json:"matched_terms"`
	Data         struct {
		ID                       string `json:"id"`
		Sku                      string `json:"sku"`
		Slug                     string `json:"slug"`
		Color                    string `json:"color"`
		Category                 string `json:"category"`
		ImageURL                 string `json:"image_url"`
		ProductType              string `json:"product_type"`
		ReleaseDate              int    `json:"release_date"`
		ReleaseDateYear          int    `json:"release_date_year"`
		RetailPriceCents         int    `json:"retail_price_cents"`
		RetailPriceCentsGbp      int    `json:"retail_price_cents_gbp"`
		RetailPriceCentsEur      int    `json:"retail_price_cents_eur"`
		RetailPriceCentsSgd      int    `json:"retail_price_cents_sgd"`
		RetailPriceCentsMyr      int    `json:"retail_price_cents_myr"`
		RetailPriceCentsCad      int    `json:"retail_price_cents_cad"`
		RetailPriceCentsCny      int    `json:"retail_price_cents_cny"`
		RetailPriceCentsAud      int    `json:"retail_price_cents_aud"`
		RetailPriceCentsKrw      int    `json:"retail_price_cents_krw"`
		RetailPriceCentsHkd      int    `json:"retail_price_cents_hkd"`
		RetailPriceCentsTwd      int    `json:"retail_price_cents_twd"`
		RetailPriceCentsJpy      int    `json:"retail_price_cents_jpy"`
		VariationID              string `json:"variation_id"`
		BoxCondition             string `json:"box_condition"`
		ProductCondition         string `json:"product_condition"`
		LowestPriceCents         int    `json:"lowest_price_cents"`
		LowestPriceCentsEur      int    `json:"lowest_price_cents_eur"`
		LowestPriceCentsMyr      int    `json:"lowest_price_cents_myr"`
		LowestPriceCentsGbp      int    `json:"lowest_price_cents_gbp"`
		LowestPriceCentsKrw      int    `json:"lowest_price_cents_krw"`
		LowestPriceCentsCad      int    `json:"lowest_price_cents_cad"`
		LowestPriceCentsTwd      int    `json:"lowest_price_cents_twd"`
		LowestPriceCentsCny      int    `json:"lowest_price_cents_cny"`
		LowestPriceCentsSgd      int    `json:"lowest_price_cents_sgd"`
		LowestPriceCentsJpy      int    `json:"lowest_price_cents_jpy"`
		LowestPriceCentsHkd      int    `json:"lowest_price_cents_hkd"`
		LowestPriceCentsAud      int    `json:"lowest_price_cents_aud"`
		CountForProductCondition int    `json:"count_for_product_condition"`
	} `json:"data"`
	Value     string   `json:"value"`
	IsSlotted bool     `json:"is_slotted"`
	Labels    struct{} `json:"labels"`
}

type response struct {
	Response struct {
		Results []GoatProduct `json:"results"`
	} `json:"response"`
}
