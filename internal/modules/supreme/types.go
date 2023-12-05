package supreme

type SupremeProductRequest struct {
	SKU string `json:"sku"`
}

type SupremeCollectionResponseProduct struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Handle      string   `json:"handle"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags"`
	ProductType string   `json:"product_type"`
	Image       string   `json:"image"`
	Images      []struct {
		Source string `json:"src"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		ID     int    `json:"id"`
		Alt    string `json:"alt"`
	} `json:"images"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Available   bool     `json:"available"`
	Collection  string   `json:"collection"`
	Options     []string `json:"options"`
	Variants    []struct {
		ID                     int           `json:"id"`
		Title                  string        `json:"title"`
		Option1                string        `json:"option1"`
		Option2                string        `json:"option2"`
		Option3                string        `json:"option3"`
		SKU                    string        `json:"sku"`
		RequiresShipping       bool          `json:"requires_shipping"`
		Taxable                bool          `json:"taxable"`
		FeaturedImage          string        `json:"featured_image"`
		Available              bool          `json:"available"`
		Name                   string        `json:"name"`
		PublicTitle            string        `json:"public_title"`
		Options                []string      `json:"options"`
		Price                  int           `json:"price"`
		Weight                 int           `json:"weight"`
		CompareAtPrice         interface{}   `json:"compare_at_price"`
		InventoryManagement    string        `json:"inventory_management"`
		Barcode                string        `json:"barcode"`
		RequiresSellingPlan    bool          `json:"requires_selling_plan"`
		SellingPlanAllocations []interface{} `json:"selling_plan_allocations"`
		QuantityRule           struct {
			Min       int         `json:"min"`
			Max       interface{} `json:"max"`
			Increment int         `json:"increment"`
		} `json:"quantity_rule"`
	} `json:"variants"`
	SelectedVariant struct {
		ID                     int           `json:"id"`
		Title                  string        `json:"title"`
		Option1                string        `json:"option1"`
		Option2                string        `json:"option2"`
		Option3                string        `json:"option3"`
		SKU                    string        `json:"sku"`
		RequiresShipping       bool          `json:"requires_shipping"`
		Taxable                bool          `json:"taxable"`
		FeaturedImage          string        `json:"featured_image"`
		Available              bool          `json:"available"`
		Name                   string        `json:"name"`
		PublicTitle            string        `json:"public_title"`
		Options                []string      `json:"options"`
		Price                  int           `json:"price"`
		Weight                 int           `json:"weight"`
		CompareAtPrice         interface{}   `json:"compare_at_price"`
		InventoryManagement    string        `json:"inventory_management"`
		Barcode                string        `json:"barcode"`
		RequiresSellingPlan    bool          `json:"requires_selling_plan"`
		SellingPlanAllocations []interface{} `json:"selling_plan_allocations"`
		QuantityRule           struct {
			Min       int         `json:"min"`
			Max       interface{} `json:"max"`
			Increment int         `json:"increment"`
		} `json:"quantity_rule"`
	} `json:"selectedVariant"`
}

type SupremeCollectionData struct {
	AllProductsCount int                                `json:"allProductsCount"`
	Products         []SupremeCollectionResponseProduct `json:"products"`
}
