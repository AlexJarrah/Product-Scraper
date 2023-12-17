package shopify

type ShopifyProduct struct {
	Product Product `json:"product"`
}
type QuantityRule struct {
	Min       int `json:"min"`
	Max       any `json:"max"`
	Increment int `json:"increment"`
}
type Variants struct {
	ID                   int64        `json:"id"`
	ProductID            int64        `json:"product_id"`
	Title                string       `json:"title"`
	Price                string       `json:"price"`
	Sku                  string       `json:"sku"`
	Position             int          `json:"position"`
	InventoryPolicy      string       `json:"inventory_policy"`
	CompareAtPrice       string       `json:"compare_at_price"`
	FulfillmentService   string       `json:"fulfillment_service"`
	InventoryManagement  string       `json:"inventory_management"`
	Option1              string       `json:"option1"`
	Option2              string       `json:"option2"`
	Option3              any          `json:"option3"`
	CreatedAt            string       `json:"created_at"`
	UpdatedAt            string       `json:"updated_at"`
	Taxable              bool         `json:"taxable"`
	Barcode              string       `json:"barcode"`
	Grams                int          `json:"grams"`
	ImageID              any          `json:"image_id"`
	Weight               float64      `json:"weight"`
	WeightUnit           string       `json:"weight_unit"`
	InventoryQuantity    int          `json:"inventory_quantity"`
	OldInventoryQuantity int          `json:"old_inventory_quantity"`
	TaxCode              string       `json:"tax_code"`
	RequiresShipping     bool         `json:"requires_shipping"`
	QuantityRule         QuantityRule `json:"quantity_rule"`
	QuantityPriceBreaks  []any        `json:"quantity_price_breaks"`
}
type Options struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}
type Images struct {
	ID         int64  `json:"id"`
	ProductID  int64  `json:"product_id"`
	Position   int    `json:"position"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Alt        string `json:"alt"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Src        string `json:"src"`
	VariantIds []any  `json:"variant_ids"`
}
type Image struct {
	ID         int64  `json:"id"`
	ProductID  int64  `json:"product_id"`
	Position   int    `json:"position"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Alt        string `json:"alt"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Src        string `json:"src"`
	VariantIds []any  `json:"variant_ids"`
}
type Product struct {
	ID             int64      `json:"id"`
	Title          string     `json:"title"`
	BodyHTML       string     `json:"body_html"`
	Vendor         string     `json:"vendor"`
	ProductType    string     `json:"product_type"`
	CreatedAt      string     `json:"created_at"`
	Handle         string     `json:"handle"`
	UpdatedAt      string     `json:"updated_at"`
	PublishedAt    string     `json:"published_at"`
	TemplateSuffix any        `json:"template_suffix"`
	PublishedScope string     `json:"published_scope"`
	Tags           string     `json:"tags"`
	Variants       []Variants `json:"variants"`
	Options        []Options  `json:"options"`
	Images         []Images   `json:"images"`
	Image          Image      `json:"image"`
}
