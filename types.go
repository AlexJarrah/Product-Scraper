package products

import "time"

type ShopifyProduct struct {
	Product Product `json:"product"`
}
type Variants struct {
	ID                   int64     `json:"id"`
	ProductID            int64     `json:"product_id"`
	Title                string    `json:"title"`
	Price                string    `json:"price"`
	Sku                  string    `json:"sku"`
	Position             int       `json:"position"`
	InventoryPolicy      string    `json:"inventory_policy"`
	CompareAtPrice       string    `json:"compare_at_price"`
	FulfillmentService   string    `json:"fulfillment_service"`
	InventoryManagement  string    `json:"inventory_management"`
	Option1              string    `json:"option1"`
	Option2              any       `json:"option2"`
	Option3              any       `json:"option3"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Taxable              bool      `json:"taxable"`
	Barcode              string    `json:"barcode"`
	Grams                int       `json:"grams"`
	ImageID              any       `json:"image_id"`
	Weight               float64   `json:"weight"`
	WeightUnit           string    `json:"weight_unit"`
	InventoryQuantity    int       `json:"inventory_quantity"`
	OldInventoryQuantity int       `json:"old_inventory_quantity"`
	RequiresShipping     bool      `json:"requires_shipping"`
}
type Options struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}
type Images struct {
	ID         int64     `json:"id"`
	ProductID  int64     `json:"product_id"`
	Position   int       `json:"position"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Alt        any       `json:"alt"`
	Width      int       `json:"width"`
	Height     int       `json:"height"`
	Src        string    `json:"src"`
	VariantIds []any     `json:"variant_ids"`
}
type Image struct {
	ID         int64     `json:"id"`
	ProductID  int64     `json:"product_id"`
	Position   int       `json:"position"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Alt        any       `json:"alt"`
	Width      int       `json:"width"`
	Height     int       `json:"height"`
	Src        string    `json:"src"`
	VariantIds []any     `json:"variant_ids"`
}
type Product struct {
	ID             int64      `json:"id"`
	Title          string     `json:"title"`
	BodyHTML       string     `json:"body_html"`
	Vendor         string     `json:"vendor"`
	ProductType    string     `json:"product_type"`
	CreatedAt      time.Time  `json:"created_at"`
	Handle         string     `json:"handle"`
	UpdatedAt      time.Time  `json:"updated_at"`
	PublishedAt    time.Time  `json:"published_at"`
	TemplateSuffix string     `json:"template_suffix"`
	PublishedScope string     `json:"published_scope"`
	Tags           string     `json:"tags"`
	Variants       []Variants `json:"variants"`
	Options        []Options  `json:"options"`
	Images         []Images   `json:"images"`
	Image          Image      `json:"image"`
}
