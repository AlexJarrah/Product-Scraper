package snipes

type SnipesProduct struct {
	Color string `json:"refinementColor"`
	Style string `json:"vendorStyle"`
	ID    string `json:"id"`
	Price struct {
		USD float32 `json:"USD"`
	} `json:"price"`
	InStock     bool   `json:"in_stock"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	AgeGroup    string `json:"ageGroup"`
	Variant     bool   `json:"variant"`
	Division    string `json:"division"`
	Master      bool   `json:"master"`
	Gender      string `json:"gender"`
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
	LongDescription string `json:"long_description"`
	URL             string `json:"url"`
	ColorCount      int    `json:"colorCount"`
}

type response struct {
	Results []struct {
		Hits []SnipesProduct `json:"hits"`
	} `json:"results"`
}
