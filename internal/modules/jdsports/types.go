package jdsports

type JDSportsProduct struct {
	Category               string   `json:"category"`
	Tags                   []string `json:"tags"`
	Department             string   `json:"department"`
	ColorHex               string   `json:"color_hex"`
	SmallImageURL          string   `json:"small_image_url"`
	Color                  string   `json:"color"`
	PriceLocalized         string   `json:"price_localized"`
	AccountID              int      `json:"account_id"`
	ProductID              string   `json:"product_id"`
	ItemID                 int      `json:"item_id"`
	Name                   string   `json:"name"`
	AffiliateLink          string   `json:"affiliate_link"`
	RetailerCategory       string   `json:"retailer_category"`
	BaseImageURL           string   `json:"base_image_url"`
	Brand                  string   `json:"brand"`
	Style                  string   `json:"style"`
	RemoteID               string   `json:"remote_id"`
	SKU                    string   `json:"sku"`
	Stocked                bool     `json:"stocked"`
	SalePrice              float32  `json:"sale_price"`
	RetailerStyle          string   `json:"retailer_style"`
	ClientOriginalImageURL string   `json:"client_original_image_url"`
	Retailer               string   `json:"retailer"`
	SalePriceLocalized     string   `json:"sale_price_localized"`
	Gender                 string   `json:"gender"`
	SKUs                   []string `json:"skus"`
	RetailerColor          string   `json:"retailer_color"`
	Price                  float32  `json:"price"`
	Pattern                string   `json:"pattern"`
}

type response struct {
	Outfits []struct {
		Items []JDSportsProduct `json:"items"`
	} `json:"outfits"`
}
