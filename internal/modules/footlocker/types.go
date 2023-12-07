package footlocker

import "time"

type FootlockerProduct struct {
	ID    string `json:"id"`
	Model struct {
		ID               string   `json:"id"`
		ModelWebKey      string   `json:"modelWebKey"`
		CompanyNumber    string   `json:"companyNumber"`
		Banner           string   `json:"banner"`
		LanguageIsoCode  string   `json:"languageIsoCode"`
		Active           bool     `json:"active"`
		Number           int      `json:"number"`
		Name             string   `json:"name"`
		Description      string   `json:"description"`
		Keywords         []any    `json:"keywords"`
		Brand            string   `json:"brand"`
		Genders          []string `json:"genders"`
		Sports           []string `json:"sports"`
		Vendor           string   `json:"vendor"`
		ProductHierarchy struct {
			ProductTypes []string `json:"productTypes"`
			Styles       []string `json:"styles"`
			SubStyles    []string `json:"subStyles"`
		} `json:"productHierarchy"`
		SizeChartID string `json:"sizeChartId"`
	} `json:"model"`
	Style struct {
		ID                string    `json:"id"`
		Sku               string    `json:"sku"`
		StyleWebKey       string    `json:"styleWebKey"`
		ModelWebKey       string    `json:"modelWebKey"`
		ModelDocumentID   string    `json:"modelDocumentId"`
		CompanyNumber     string    `json:"companyNumber"`
		Banner            string    `json:"banner"`
		LanguageIsoCode   string    `json:"languageIsoCode"`
		Active            bool      `json:"active"`
		Description       string    `json:"description"`
		Color             string    `json:"color"`
		PrimaryColor      string    `json:"primaryColor"`
		SecondaryColors   []string  `json:"secondaryColors"`
		Width             string    `json:"width"`
		LeagueName        string    `json:"leagueName"`
		PlayerName        string    `json:"playerName"`
		TeamName          string    `json:"teamName"`
		FitVariant        string    `json:"fitVariant"`
		Keywords          []string  `json:"keywords"`
		ProductDesignator string    `json:"productDesignator"`
		NewArrivalDate    time.Time `json:"newArrivalDate"`
		AgeBuckets        []any     `json:"ageBuckets"`
		Price             struct {
			CurrencyIso                  any     `json:"currencyIso"`
			ListPrice                    float64 `json:"listPrice"`
			SalePrice                    float64 `json:"salePrice"`
			VendorShippingPrice          any     `json:"vendorShippingPrice"`
			FormattedListPrice           string  `json:"formattedListPrice"`
			FormattedSalePrice           string  `json:"formattedSalePrice"`
			FormattedVendorShippingPrice string  `json:"formattedVendorShippingPrice"`
			PriceRange                   string  `json:"priceRange"`
			TopSalesAmount               float64 `json:"topSalesAmount"`
			TaxClassificationCode        string  `json:"taxClassificationCode"`
		} `json:"price"`
		PriceNet             any `json:"priceNet"`
		FlagsAndRestrictions struct {
			DefaultStyle            bool `json:"defaultStyle"`
			NewProduct              bool `json:"newProduct"`
			SaleProduct             bool `json:"saleProduct"`
			ExcludedFromDiscount    bool `json:"excludedFromDiscount"`
			MapEnabled              bool `json:"mapEnabled"`
			FreeShipping            bool `json:"freeShipping"`
			RecaptchaOn             bool `json:"recaptchaOn"`
			ShipToAndFromStore      bool `json:"shipToAndFromStore"`
			HasShippingRestrictions bool `json:"hasShippingRestrictions"`
			HasVendorShippingPrice  bool `json:"hasVendorShippingPrice"`
			CanBePaidByKlarna       any  `json:"canBePaidByKlarna"`
		} `json:"flagsAndRestrictions"`
		LaunchAttributes struct {
			LaunchProduct                     bool   `json:"launchProduct"`
			LaunchType                        string `json:"launchType"`
			WebOnlyLaunchMsg                  string `json:"webOnlyLaunchMsg"`
			WebOnlyLaunch                     bool   `json:"webOnlyLaunch"`
			LaunchDate                        any    `json:"launchDate"`
			LaunchDisplayCounterEnabled       bool   `json:"launchDisplayCounterEnabled"`
			LaunchDisplayCounterKickStartTime any    `json:"launchDisplayCounterKickStartTime"`
		} `json:"launchAttributes"`
		GiftCardDenominations any `json:"giftCardDenominations"`
		EligiblePaymentTypes  struct {
			CreditCard bool `json:"creditCard"`
			GiftCard   bool `json:"giftCard"`
			PayPal     bool `json:"payPal"`
			Klarna     bool `json:"klarna"`
			ApplePay   bool `json:"applePay"`
			GooglePay  bool `json:"googlePay"`
			PayBright  bool `json:"payBright"`
			ClearPay   any  `json:"clearPay"`
			IdealPay   any  `json:"idealPay"`
			Sofort     any  `json:"sofort"`
		} `json:"eligiblePaymentTypes"`
		VendorAttributes struct {
			SupplierSkus []string `json:"supplierSkus"`
		} `json:"vendorAttributes"`
		ImageURL struct {
			Base     string   `json:"base"`
			ImageSku string   `json:"imageSku"`
			Variants []string `json:"variants"`
		} `json:"imageUrl"`
	} `json:"style"`
	Inventory struct {
		InventoryAvailable          bool  `json:"inventoryAvailable"`
		StoreInventoryAvailable     bool  `json:"storeInventoryAvailable"`
		WarehouseInventoryAvailable bool  `json:"warehouseInventoryAvailable"`
		DropshipInventoryAvailable  bool  `json:"dropshipInventoryAvailable"`
		InventoryAvailableLocations []any `json:"inventoryAvailableLocations"`
		PreSell                     any   `json:"preSell"`
		BackOrder                   any   `json:"backOrder"`
		PurchaseOrderDate           any   `json:"purchaseOrderDate"`
	} `json:"inventory"`
	Sizes []struct {
		ID              string `json:"id"`
		ProductWebKey   string `json:"productWebKey"`
		StyleDocumentID string `json:"styleDocumentId"`
		ModelDocumentID string `json:"modelDocumentId"`
		CompanyNumber   string `json:"companyNumber"`
		Banner          string `json:"banner"`
		LanguageIsoCode string `json:"languageIsoCode"`
		Active          bool   `json:"active"`
		ProductNumber   string `json:"productNumber"`
		Size            string `json:"size"`
		StrippedSize    string `json:"strippedSize"`
		SizeVariants    any    `json:"sizeVariants"`
		Upc             string `json:"upc"`
		StoreUpc        string `json:"storeUpc"`
		StoreSku        string `json:"storeSku"`
		Price           struct {
			CurrencyIso                  any     `json:"currencyIso"`
			ListPrice                    float64 `json:"listPrice"`
			SalePrice                    float64 `json:"salePrice"`
			VendorShippingPrice          float64 `json:"vendorShippingPrice"`
			FormattedListPrice           string  `json:"formattedListPrice"`
			FormattedSalePrice           string  `json:"formattedSalePrice"`
			FormattedVendorShippingPrice string  `json:"formattedVendorShippingPrice"`
			PriceRange                   string  `json:"priceRange"`
			TopSalesAmount               any     `json:"topSalesAmount"`
			TaxClassificationCode        any     `json:"taxClassificationCode"`
		} `json:"price"`
		PriceNet  any `json:"priceNet"`
		Inventory struct {
			InventoryAvailable          bool  `json:"inventoryAvailable"`
			StoreInventoryAvailable     bool  `json:"storeInventoryAvailable"`
			WarehouseInventoryAvailable bool  `json:"warehouseInventoryAvailable"`
			DropshipInventoryAvailable  bool  `json:"dropshipInventoryAvailable"`
			InventoryAvailableLocations []any `json:"inventoryAvailableLocations"`
			PreSell                     bool  `json:"preSell"`
			BackOrder                   bool  `json:"backOrder"`
			PurchaseOrderDate           any   `json:"purchaseOrderDate"`
		} `json:"inventory"`
	} `json:"sizes"`
	StyleVariants []struct {
		Active        bool   `json:"active"`
		StyleWebKey   string `json:"styleWebKey"`
		Sku           string `json:"sku"`
		Upc           string `json:"upc"`
		StoreUpc      string `json:"storeUpc"`
		ProductWebKey string `json:"productWebKey"`
		Style         string `json:"style"`
		Color         string `json:"color"`
		Description   string `json:"description"`
		PlayerName    string `json:"playerName"`
		ImageSku      string `json:"imageSku"`
		Size          string `json:"size"`
		AgeBuckets    any    `json:"ageBuckets"`
		Price         struct {
			CurrencyIso                  any     `json:"currencyIso"`
			ListPrice                    float64 `json:"listPrice"`
			SalePrice                    float64 `json:"salePrice"`
			VendorShippingPrice          float64 `json:"vendorShippingPrice"`
			FormattedListPrice           string  `json:"formattedListPrice"`
			FormattedSalePrice           string  `json:"formattedSalePrice"`
			FormattedVendorShippingPrice string  `json:"formattedVendorShippingPrice"`
			PriceRange                   string  `json:"priceRange"`
			TopSalesAmount               any     `json:"topSalesAmount"`
			TaxClassificationCode        any     `json:"taxClassificationCode"`
		} `json:"price"`
		PriceNet  any `json:"priceNet"`
		Inventory struct {
			InventoryAvailable          bool  `json:"inventoryAvailable"`
			StoreInventoryAvailable     bool  `json:"storeInventoryAvailable"`
			WarehouseInventoryAvailable bool  `json:"warehouseInventoryAvailable"`
			DropshipInventoryAvailable  bool  `json:"dropshipInventoryAvailable"`
			InventoryAvailableLocations []any `json:"inventoryAvailableLocations"`
			PreSell                     bool  `json:"preSell"`
			BackOrder                   bool  `json:"backOrder"`
			PurchaseOrderDate           any   `json:"purchaseOrderDate"`
		} `json:"inventory"`
	} `json:"styleVariants"`
	SizeChart struct {
		ID               string `json:"id"`
		SizeChartGridMap []struct {
			Label string   `json:"label"`
			Sizes []string `json:"sizes"`
		} `json:"sizeChartGridMap"`
		SizeChartTipTx string `json:"sizeChartTipTx"`
		SizeChartImage string `json:"sizeChartImage"`
	} `json:"sizeChart"`
}

type response FootlockerProduct
