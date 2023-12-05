package jdsports

type JDSportsRequest struct {
	SKU string `json:"sku"`
}

type JDSportsData struct {
	Outfits []struct {
		// Tags  []string `json:"tags"`
		PrimaryStylitics   int    `json:"primary_stylitics_item_id"`
		RelativeBoost      string `json:"relative_boost"`
		AccountID          int    `json:"account_id"`
		Username           string `json:"username"`
		PrimaryRequestedID string `json:"primary_requested_id"`
		BaseImageURL       string `json:"base_image_url"`
		PrimaryRemoteID    string `json:"primary_remote_id"`
		UpdatedAt          string `json:"updated_at"`
		// CoverImageURL      string `json:"cover_image_url"`
		ID              int    `json:"id"`
		AccountUsername string `json:"account_username"`
		Items           []struct {
			// Promotions []string `json:"promotions"`
			Category   string   `json:"category"`
			Tags       []string `json:"tags"`
			Department string   `json:"department"`
			ColorHex   string   `json:"color_hex"`
			// OtherClientItemIDs []int    `json:"other_client_item_ids"`
			SmallImageURL  string `json:"small_image_url"`
			Color          string `json:"color"`
			PriceLocalized string `json:"price_localized"`
			AccountID      int    `json:"account_id"`
			ProductID      string `json:"product_id"`
			Coords         struct {
				X      int `json:"x"`
				Y      int `json:"y"`
				Z      int `json:"z"`
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"coords"`
			ItemID           int    `json:"item_id"`
			Name             string `json:"name"`
			AffiliateLink    string `json:"affiliate_link"`
			Username         string `json:"username"`
			RetailerCategory string `json:"retailer_category"`
			BaseImageURL     string `json:"base_image_url"`
			Brand            string `json:"brand"`
			Style            string `json:"style"`
			RemoteID         string `json:"remote_id"`
			// SKU              string `json:"sku"`
			Stocked   bool    `json:"stocked"`
			SalePrice float32 `json:"sale_price"`
			// RetailerStyle string `json:"retailer_style"`
			ClientOriginalImageURL string `json:"client_original_image_url"`
			AccountUsername        string `json:"account_username"`
			// Retailer               string `json:"retailer"`
			SalePriceLocalized string   `json:"sale_price_localized"`
			Gender             string   `json:"gender"`
			SKUs               []string `json:"skus"`
			RetailerColor      string   `json:"retailer_color"`
			Price              float32  `json:"price"`
			Pattern            string   `json:"pattern"`
		} `json:"items"`
	} `json:"outfits"`
}
