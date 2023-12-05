package adidas

type AdidasRequest struct {
	SKUs []string `json:"skus"`
}

type AdidasData []struct {
	ID          string `json:"id"`
	ProductType string `json:"product_type"`
	ModelNumber string `json:"model_number"`
	Name        string `json:"name"`
	Metadata    struct {
		Canonical   string `json:"canonical"`
		Description string `json:"description"`
		Keywords    string `json:"keywords"`
		PageTitle   string `json:"page_title"`
		SiteName    string `json:"site_name"`
	} `json:"meta_data"`
	ViewList []struct {
		Type     string `json:"type"`
		Source   string `json:"source"`
		ImageURL string `json:"image_url"`
		Metadata struct {
			AssetUsage    []string `json:"asset_usage"`
			AssetCategory string   `json:"asset_category"`
			View          string   `json:"view"`
			UsageTerms    string   `json:"usage_terms"`
			SortOrder     string   `json:"sort_order"`
			Subjects      []string `json:"subjects"`
		} `json:"metadata"`
	} `json:"view_list"`
	AttributeList struct {
		Sale                     bool                   `json:"sale"`
		Brand                    string                 `json:"brand"`
		Color                    string                 `json:"color"`
		Gender                   string                 `json:"gender"`
		Outlet                   bool                   `json:"outlet"`
		Sport                    []string               `json:"sport"`
		Category                 string                 `json:"category"`
		SizePage                 string                 `json:"size_page"`
		SportIDs                 []string               `json:"sport_ids"`
		SportSub                 []string               `json:"sportSub"`
		BestForIDs               []string               `json:"best_for_ids"`
		ProductFit               []string               `json:"productfit"`
		Collection               []string               `json:"collection"`
		SearchColor              string                 `json:"search_color"`
		BaseMaterial             []string               `json:"base_material"`
		ProductType              []string               `json:"productType"`
		MidsoleOffset            string                 `json:"midsole_offset"`
		ProductFitIDs            []string               `json:"product_fit_ids"`
		Personalizable           bool                   `json:"personalizable"`
		ToeStackHeight           string                 `json:"toe_stack_height"`
		HeelStackHeight          string                 `json:"heel_stack_height"`
		IsCnCRestricted          bool                   `json:"isCnCRestricted"`
		KeyCategoryCode          string                 `json:"key_category_code"`
		MandatoryPersonalization bool                   `json:"mandatory_personalization"`
		Customizable             bool                   `json:"customizable"`
		BadgeStyle               string                 `json:"badge_style"`
		BadgeText                string                 `json:"badge_text"`
		SearchColorRaw           string                 `json:"search_color_raw"`
		IsOrderable              bool                   `json:"is_orderable"`
		IsWaitingRoomProduct     bool                   `json:"isWaitingRoomProduct"`
		IsInPreview              bool                   `json:"isInPreview"`
		SpecialLaunch            bool                   `json:"specialLaunch"`
		SpecialLaunchType        string                 `json:"special_launch_type"`
		SizeTypes                map[string]interface{} `json:"sizeTypes"`
		IsFlash                  bool                   `json:"is_flash"`
		IsMadeToBeRemade         bool                   `json:"is_made_to_be_remade"`
		ProductSizingCategory    string                 `json:"product_sizing_category"`
		SizeChartID              string                 `json:"size_chart_id"`
		SizeChartLink            string                 `json:"size_chart_link"`
	} `json:"attribute_list"`
	BreadcrumbList []struct {
		Text string `json:"text"`
		Link string `json:"link"`
	} `json:"breadcrumb_list"`
	Callouts struct {
		CalloutTopStack []struct {
			ID       string `json:"id"`
			SubTitle string `json:"sub_title"`
			Body     string `json:"body"`
		} `json:"callout_top_stack"`
	} `json:"callouts"`
	PricingInformation struct {
		CurrentPrice       int `json:"currentPrice"`
		StandardPrice      int `json:"standard_price"`
		StandardPriceNoVAT int `json:"standard_price_no_vat"`
	} `json:"pricing_information"`
	TaxClassID         string `json:"tax_class_id"`
	ProductDescription struct {
		Title                string   `json:"title"`
		Text                 string   `json:"text"`
		Subtitle             string   `json:"subtitle"`
		USPs                 []string `json:"usps"`
		WashCareInstructions struct {
			CareInstructions []string `json:"care_instructions"`
		} `json:"wash_care_instructions"`
		DescriptionAssets struct {
			ImageURL  string `json:"image_url"`
			VideoURL  string `json:"video_url"`
			PosterURL string `json:"poster_url"`
		} `json:"description_assets"`
	} `json:"product_description"`
	RecommendationsEnabled bool `json:"recommendationsEnabled"`
	ProductLinkList        []struct {
		Type                            string `json:"type"`
		ProductID                       string `json:"productId"`
		Name                            string `json:"name"`
		URL                             string `json:"url"`
		Image                           string `json:"image"`
		AltImage                        string `json:"altImage"`
		DynamicBackgroundImage          string `json:"dynamic_background_image"`
		ConfirmedDynamicBackgroundImage string `json:"confirmed_dynamic_background_image"`
		PricingInformation              struct {
			StandardPrice int `json:"standard_price"`
		} `json:"pricing_information"`
		BadgeStyle   string `json:"badge_style"`
		BadgeText    string `json:"badge_text"`
		SearchColor  string `json:"search_color"`
		DefaultColor string `json:"default_color"`
		// TeamKits      []interface{} `json:"team_kits"`
		Source        string `json:"source"`
		AvailableSKUs int    `json:"available_skus"`
	} `json:"product_link_list"`
	VariationList []struct {
		SKU  string `json:"sku"`
		Size string `json:"size"`
	} `json:"variation_list"`
}
