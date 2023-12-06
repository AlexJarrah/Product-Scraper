package adidas

import "time"

type AdidasProduct struct {
	ID          string `json:"id"`
	ProductType string `json:"product_type"`
	ModelNumber string `json:"model_number"`
	Name        string `json:"name"`
	MetaData    struct {
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
			ImageStyle    string   `json:"imageStyle"`
			View          string   `json:"view"`
			UsageTerms    string   `json:"usageTerms"`
			SortOrder     string   `json:"sortOrder"`
			Subjects      []any    `json:"subjects"`
		} `json:"metadata"`
		FileRevisionDate time.Time `json:"file_revision_date"`
	} `json:"view_list"`
	DynamicBackgroundAssets []struct {
		Type     string `json:"type"`
		Source   string `json:"source"`
		ImageURL string `json:"image_url"`
		Metadata struct {
			AssetUsage    []string `json:"asset_usage"`
			AssetCategory string   `json:"asset_category"`
			ImageStyle    string   `json:"image_style"`
			View          string   `json:"view"`
			UsageTerms    string   `json:"usage_terms"`
			SortOrder     string   `json:"sort_order"`
			Subjects      []any    `json:"subjects"`
		} `json:"metadata"`
		FileRevisionDate time.Time `json:"file_revision_date"`
	} `json:"dynamic_background_assets"`
	ConfirmedDynamicBackgroundAssets []struct {
		Type     string `json:"type"`
		Source   string `json:"source"`
		ImageURL string `json:"image_url"`
		Metadata struct {
			AssetUsage    []string `json:"asset_usage"`
			AssetCategory string   `json:"asset_category"`
			ImageStyle    string   `json:"image_style"`
			View          string   `json:"view"`
			UsageTerms    string   `json:"usage_terms"`
			SortOrder     string   `json:"sort_order"`
			Subjects      []any    `json:"subjects"`
		} `json:"metadata"`
		FileRevisionDate time.Time `json:"file_revision_date"`
	} `json:"confirmed_dynamic_background_assets"`
	AttributeList struct {
		Sale                     bool     `json:"sale"`
		Brand                    string   `json:"brand"`
		Color                    string   `json:"color"`
		Gender                   string   `json:"gender"`
		Outlet                   bool     `json:"outlet"`
		Sport                    []string `json:"sport"`
		Closure                  []string `json:"closure"`
		Surface                  []string `json:"surface"`
		Category                 string   `json:"category"`
		SizePage                 string   `json:"size_page"`
		SportIds                 []string `json:"sport_ids"`
		SportSub                 []string `json:"sportSub"`
		CategoryID               []string `json:"category_id"`
		Productfit               []string `json:"productfit"`
		SurfaceIds               []string `json:"surface_ids"`
		Collection               []string `json:"collection"`
		SearchColor              string   `json:"search_color"`
		BaseMaterial             []string `json:"base_material"`
		ProductType              []string `json:"productType"`
		MidsoleOffset            string   `json:"midsole_offset"`
		ProductFitIds            []string `json:"product_fit_ids"`
		Personalizable           bool     `json:"personalizable"`
		ToeStackHeight           string   `json:"toe_stack_height"`
		HeelStackHeight          string   `json:"heel_stack_height"`
		IsCnCRestricted          bool     `json:"isCnCRestricted"`
		KeyCategoryCode          string   `json:"key_category_code"`
		ProductLineStyle         []string `json:"productLineStyle"`
		MandatoryPersonalization bool     `json:"mandatory_personalization"`
		Customizable             bool     `json:"customizable"`
		BadgeStyle               string   `json:"badge_style"`
		BadgeText                string   `json:"badge_text"`
		SearchColorRaw           string   `json:"search_color_raw"`
		IsOrderable              bool     `json:"is_orderable"`
		IsWaitingRoomProduct     bool     `json:"isWaitingRoomProduct"`
		IsInPreview              bool     `json:"isInPreview"`
		SpecialLaunch            bool     `json:"specialLaunch"`
		SpecialLaunchType        string   `json:"special_launch_type"`
		SizeTypes                struct {
		} `json:"sizeTypes"`
		IsFlash               bool      `json:"is_flash"`
		IsMadeToBeRemade      bool      `json:"is_made_to_be_remade"`
		ComingSoonSignup      bool      `json:"coming_soon_signup"`
		PreviewTo             time.Time `json:"preview_to"`
		ProductSizingCategory string    `json:"product_sizing_category"`
		SizeChartID           string    `json:"size_chart_id"`
		SizeChartLink         string    `json:"size_chart_link"`
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
		StandardPriceNoVat int `json:"standard_price_no_vat"`
	} `json:"pricing_information"`
	TaxClassID         string `json:"tax_class_id"`
	ProductDescription struct {
		Title                string   `json:"title"`
		Text                 string   `json:"text"`
		Subtitle             string   `json:"subtitle"`
		Usps                 []string `json:"usps"`
		WashCareInstructions struct {
			CareInstructions []any `json:"care_instructions"`
		} `json:"wash_care_instructions"`
		DescriptionAssets struct {
			ImageURL  string `json:"image_url"`
			VideoURL  any    `json:"video_url"`
			PosterURL any    `json:"poster_url"`
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
		DynamicBackgroundImage          string `json:"dynamic_background_image,omitempty"`
		ConfirmedDynamicBackgroundImage string `json:"confirmed_dynamic_background_image,omitempty"`
		PricingInformation              struct {
			StandardPrice int `json:"standard_price"`
		} `json:"pricing_information"`
		BadgeStyle    string `json:"badge_style"`
		BadgeText     string `json:"badge_text"`
		SearchColor   string `json:"search_color"`
		DefaultColor  string `json:"default_color"`
		TeamKits      []any  `json:"team_kits"`
		Source        string `json:"source"`
		AvailableSkus int    `json:"available_skus"`
	} `json:"product_link_list"`
	VariationList []struct {
		Sku  string `json:"sku"`
		Size string `json:"size"`
	} `json:"variation_list"`
}

type response []AdidasProduct
