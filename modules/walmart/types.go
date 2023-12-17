package walmart

import "time"

type WalmartProductRequest struct {
	SKU string `json:"sku"`
}

type WalmartProductData struct {
	FulfillmentTypeClassification any  `json:"fulfillmentTypeClassification"`
	Ebooks                        any  `json:"ebooks"`
	BlitzItem                     any  `json:"blitzItem"`
	GiftingEligibility            bool `json:"giftingEligibility"`
	ShipAsIs                      bool `json:"shipAsIs"`
	Subscription                  struct {
		ShowSubscriptionModule   bool   `json:"showSubscriptionModule"`
		SubscriptionEligible     bool   `json:"subscriptionEligible"`
		SubscriptionTransactable bool   `json:"subscriptionTransactable"`
		SubscriptionSubmessage   string `json:"subscriptionSubmessage"`
	} `json:"subscription"`
	SubscriptionEligible   bool    `json:"subscriptionEligible"`
	ShowSubscriptionModule bool    `json:"showSubscriptionModule"`
	BuyBoxSuppression      bool    `json:"buyBoxSuppression"`
	Discounts              any     `json:"discounts"`
	Rewards                any     `json:"rewards"`
	PromoDiscount          any     `json:"promoDiscount"`
	ShowFulfillmentLink    bool    `json:"showFulfillmentLink"`
	AdditionalOfferCount   int     `json:"additionalOfferCount"`
	LegalRestriction       bool    `json:"legalRestriction"`
	ShippingRestriction    bool    `json:"shippingRestriction"`
	AvailabilityStatus     string  `json:"availabilityStatus"`
	AverageRating          float64 `json:"averageRating"`
	AssociatedBundleID     any     `json:"associatedBundleId"`
	SuppressReviews        bool    `json:"suppressReviews"`
	Brand                  string  `json:"brand"`
	BrandURL               string  `json:"brandUrl"`
	Badges                 struct {
		Flags []struct {
			Typename string `json:"__typename"`
			ID       string `json:"id"`
			Text     string `json:"text"`
			Key      string `json:"key"`
			Query    any    `json:"query"`
			Type     string `json:"type"`
			StyleID  any    `json:"styleId"`
		} `json:"flags"`
		Labels any   `json:"labels"`
		Tags   []any `json:"tags"`
		Groups []any `json:"groups"`
	} `json:"badges"`
	RhPath                string `json:"rhPath"`
	AaiaBrandID           any    `json:"aaiaBrandId"`
	ManufacturerProductID string `json:"manufacturerProductId"`
	ProductTypeID         string `json:"productTypeId"`
	TireSize              any    `json:"tireSize"`
	TireLoadIndex         any    `json:"tireLoadIndex"`
	TireSpeedRating       any    `json:"tireSpeedRating"`
	Viscosity             any    `json:"viscosity"`
	MotorOilType          any    `json:"motorOilType"`
	Model                 string `json:"model"`
	BuyNowEligible        bool   `json:"buyNowEligible"`
	PetRxEligible         any    `json:"petRxEligible"`
	PetRx                 struct {
		Eligible       bool `json:"eligible"`
		SingleDispense any  `json:"singleDispense"`
	} `json:"petRx"`
	Vision struct {
		VisionCenterApproved bool   `json:"visionCenterApproved"`
		AgeGroup             string `json:"ageGroup"`
	} `json:"vision"`
	EarlyAccessEvent  bool `json:"earlyAccessEvent"`
	IsEarlyAccessItem bool `json:"isEarlyAccessItem"`
	AnnualEvent       bool `json:"annualEvent"`
	EventAttributes   struct {
		PriceFlip  bool `json:"priceFlip"`
		SpecialBuy bool `json:"specialBuy"`
	} `json:"eventAttributes"`
	IsWplusMember    bool `json:"isWplusMember"`
	ShowBuyWithWplus bool `json:"showBuyWithWplus"`
	PreOrder         struct {
		IsPreOrder bool `json:"isPreOrder"`
	} `json:"preOrder"`
	OzarkAttributes struct {
		ShippingPromise any  `json:"shippingPromise"`
		ExactAddress    bool `json:"exactAddress"`
	} `json:"ozarkAttributes"`
	CanonicalURL        string  `json:"canonicalUrl"`
	CatalogSellerID     int     `json:"catalogSellerId"`
	SellerReviewCount   int     `json:"sellerReviewCount"`
	SellerAverageRating float64 `json:"sellerAverageRating"`
	Category            struct {
		CategoryPathID string `json:"categoryPathId"`
		Path           []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"path"`
	} `json:"category"`
	ClassType                  string `json:"classType"`
	ClassID                    string `json:"classId"`
	FulfillmentTitle           string `json:"fulfillmentTitle"`
	ShortDescription           string `json:"shortDescription"`
	FulfillmentType            string `json:"fulfillmentType"`
	FulfillmentBadge           any    `json:"fulfillmentBadge"`
	CheckStoreAvailabilityATC  bool   `json:"checkStoreAvailabilityATC"`
	CheckAvailabilityGlobalDFS bool   `json:"checkAvailabilityGlobalDFS"`
	FulfillmentLabel           []struct {
		CheckStoreAvailability any    `json:"checkStoreAvailability"`
		WPlusFulfillmentText   any    `json:"wPlusFulfillmentText"`
		Message                string `json:"message"`
		ShippingText           string `json:"shippingText"`
		FulfillmentText        string `json:"fulfillmentText"`
		LocationText           string `json:"locationText"`
		FulfillmentMethod      string `json:"fulfillmentMethod"`
		AddressEligibility     bool   `json:"addressEligibility"`
		FulfillmentType        string `json:"fulfillmentType"`
		PostalCode             string `json:"postalCode"`
	} `json:"fulfillmentLabel"`
	FulfillmentOptions []struct {
		Typename          string `json:"__typename"`
		Type              string `json:"type"`
		Selected          bool   `json:"selected"`
		Intent            bool   `json:"intent"`
		AvailableQuantity int    `json:"availableQuantity"`
		MaxOrderQuantity  int    `json:"maxOrderQuantity"`
		OrderLimit        int    `json:"orderLimit"`
		SpeedDetails      struct {
			FulfillmentBadge any       `json:"fulfillmentBadge"`
			DeliveryDate     time.Time `json:"deliveryDate"`
			FulfillmentPrice any       `json:"fulfillmentPrice"`
			FreeFulfillment  bool      `json:"freeFulfillment"`
			WPlusEligible    bool      `json:"wPlusEligible"`
		} `json:"speedDetails"`
		LocationText           string `json:"locationText"`
		AvailabilityStatus     string `json:"availabilityStatus"`
		SubscriptionSubmessage string `json:"subscriptionSubmessage"`
		InventoryStatus        any    `json:"inventoryStatus,omitempty"`
		Restricted             bool   `json:"restricted,omitempty"`
		ProductLocation        any    `json:"productLocation,omitempty"`
		StoreName              any    `json:"storeName,omitempty"`
		CheckStoreAvailability bool   `json:"checkStoreAvailability,omitempty"`
	} `json:"fulfillmentOptions"`
	HasSellerBadge  bool   `json:"hasSellerBadge"`
	HasCarePlans    bool   `json:"hasCarePlans"`
	HasHomeServices any    `json:"hasHomeServices"`
	ItemType        any    `json:"itemType"`
	ID              string `json:"id"`
	PrimaryUsItemID any    `json:"primaryUsItemId"`
	ConditionType   string `json:"conditionType"`
	ImageInfo       struct {
		AllImages []struct {
			ID       string `json:"id"`
			URL      string `json:"url"`
			Zoomable bool   `json:"zoomable"`
		} `json:"allImages"`
		ThumbnailURL string `json:"thumbnailUrl"`
	} `json:"imageInfo"`
	Location struct {
		PostalCode          string   `json:"postalCode"`
		StateOrProvinceCode string   `json:"stateOrProvinceCode"`
		City                string   `json:"city"`
		StoreIds            []string `json:"storeIds"`
		AddressID           any      `json:"addressId"`
		Intent              string   `json:"intent"`
		MpPickupLocation    any      `json:"mpPickupLocation"`
		PickupLocation      struct {
			StoreID       string `json:"storeId"`
			AccessPointID any    `json:"accessPointId"`
			AccessType    any    `json:"accessType"`
		} `json:"pickupLocation"`
	} `json:"location"`
	Name            string `json:"name"`
	Personalizable  bool   `json:"personalizable"`
	ExternalInfo    any    `json:"externalInfo"`
	NumberOfReviews int    `json:"numberOfReviews"`
	OrderMinLimit   int    `json:"orderMinLimit"`
	OrderLimit      int    `json:"orderLimit"`
	WeightIncrement int    `json:"weightIncrement"`
	OfferID         string `json:"offerId"`
	OfferType       string `json:"offerType"`
	PriceInfo       struct {
		PriceDisplayCodes struct {
			Clearance                    any `json:"clearance"`
			EligibleForAssociateDiscount any `json:"eligibleForAssociateDiscount"`
			FinalCostByWeight            any `json:"finalCostByWeight"`
			PriceDisplayCondition        any `json:"priceDisplayCondition"`
			ReducedPrice                 any `json:"reducedPrice"`
			Rollback                     any `json:"rollback"`
			SubmapType                   any `json:"submapType"`
		} `json:"priceDisplayCodes"`
		CurrentPrice struct {
			Price              float64 `json:"price"`
			PriceString        string  `json:"priceString"`
			VariantPriceString string  `json:"variantPriceString"`
			CurrencyUnit       string  `json:"currencyUnit"`
			BestValue          any     `json:"bestValue"`
			PriceDisplay       string  `json:"priceDisplay"`
		} `json:"currentPrice"`
		WasPrice            any  `json:"wasPrice"`
		ComparisonPrice     any  `json:"comparisonPrice"`
		UnitPrice           any  `json:"unitPrice"`
		Savings             any  `json:"savings"`
		SavingsAmount       any  `json:"savingsAmount"`
		SecondaryOfferBoost any  `json:"secondaryOfferBoost"`
		ShipPrice           any  `json:"shipPrice"`
		IsPriceReduced      bool `json:"isPriceReduced"`
		PriceReducedDisplay any  `json:"priceReducedDisplay"`
		SubscriptionPrice   any  `json:"subscriptionPrice"`
		PriceRange          struct {
			MinPrice      any `json:"minPrice"`
			MaxPrice      any `json:"maxPrice"`
			PriceString   any `json:"priceString"`
			CurrencyUnit  any `json:"currencyUnit"`
			Denominations any `json:"denominations"`
		} `json:"priceRange"`
		ListPrice             any `json:"listPrice"`
		CapType               any `json:"capType"`
		WalmartFundedAmount   any `json:"walmartFundedAmount"`
		WPlusEarlyAccessPrice any `json:"wPlusEarlyAccessPrice"`
	} `json:"priceInfo"`
	ReturnPolicy struct {
		Returnable   bool `json:"returnable"`
		FreeReturns  bool `json:"freeReturns"`
		ReturnWindow struct {
			Value    int    `json:"value"`
			UnitType string `json:"unitType"`
		} `json:"returnWindow"`
		ReturnPolicyText      string `json:"returnPolicyText"`
		ReturnPolicyCondition any    `json:"returnPolicyCondition"`
		HolidayReturnEnabled  bool   `json:"holidayReturnEnabled"`
	} `json:"returnPolicy"`
	FsaEligibleInd      bool   `json:"fsaEligibleInd"`
	SellerID            string `json:"sellerId"`
	SellerName          string `json:"sellerName"`
	SellerDisplayName   string `json:"sellerDisplayName"`
	SecondaryOfferPrice struct {
		CurrentPrice struct {
			PriceType   any     `json:"priceType"`
			PriceString string  `json:"priceString"`
			Price       float64 `json:"price"`
		} `json:"currentPrice"`
	} `json:"secondaryOfferPrice"`
	SemStoreData   any `json:"semStoreData"`
	ShippingOption struct {
		AvailabilityStatus string    `json:"availabilityStatus"`
		SLATier            string    `json:"slaTier"`
		DeliveryDate       time.Time `json:"deliveryDate"`
		MaxDeliveryDate    any       `json:"maxDeliveryDate"`
		ShipMethod         string    `json:"shipMethod"`
		ShipPrice          struct {
			Price              int    `json:"price"`
			PriceString        string `json:"priceString"`
			VariantPriceString any    `json:"variantPriceString"`
			CurrencyUnit       any    `json:"currencyUnit"`
			BestValue          any    `json:"bestValue"`
		} `json:"shipPrice"`
		InternationalShipping any `json:"internationalShipping"`
	} `json:"shippingOption"`
	Type         string `json:"type"`
	PickupOption struct {
		SLATier            any   `json:"slaTier"`
		AccessTypes        []any `json:"accessTypes"`
		AvailabilityStatus any   `json:"availabilityStatus"`
		StoreName          any   `json:"storeName"`
		StoreID            any   `json:"storeId"`
	} `json:"pickupOption"`
	SalesUnit                 string `json:"salesUnit"`
	UsItemID                  string `json:"usItemId"`
	TransactableOfferCount    int    `json:"transactableOfferCount"`
	MaxRewardAmongAllVariants any    `json:"maxRewardAmongAllVariants"`
	VariantCriteria           []any  `json:"variantCriteria"`
	Variants                  []any  `json:"variants"`
	GroupMetaData             any    `json:"groupMetaData"`
	Upc                       string `json:"upc"`
	WfsEnabled                bool   `json:"wfsEnabled"`
	WfsProviderName           any    `json:"wfsProviderName"`
	SellerType                string `json:"sellerType"`
	IronbankCategory          string `json:"ironbankCategory"`
	SnapEligible              bool   `json:"snapEligible"`
	PromoData                 []any  `json:"promoData"`
	ShowAddOnServices         bool   `json:"showAddOnServices"`
	AddOnServices             []struct {
		ServiceType     string `json:"serviceType"`
		ServiceTitle    string `json:"serviceTitle"`
		ServiceSubTitle string `json:"serviceSubTitle"`
		Groups          []struct {
			GroupType            string `json:"groupType"`
			GroupTitle           string `json:"groupTitle"`
			AssetURL             string `json:"assetUrl"`
			ShortDescription     string `json:"shortDescription"`
			UnavailabilityReason any    `json:"unavailabilityReason"`
			NearByStores         any    `json:"nearByStores"`
			Services             []struct {
				Name                string `json:"name"`
				DisplayName         string `json:"displayName"`
				OfferID             string `json:"offerId"`
				UsItemID            string `json:"usItemId"`
				SelectedDisplayName string `json:"selectedDisplayName"`
				ServiceMetaData     any    `json:"serviceMetaData"`
				CurrentPrice        struct {
					Price       int    `json:"price"`
					PriceString string `json:"priceString"`
				} `json:"currentPrice"`
				GiftEligible bool `json:"giftEligible"`
			} `json:"services"`
		} `json:"groups"`
	} `json:"addOnServices"`
	ProductLocation any `json:"productLocation"`
	ZeekitData      any `json:"zeekitData"`
	ExperienceType  any `json:"experienceType"`
	Redirect        struct {
		ReplacedByItemID         string `json:"replacedByItemId"`
		ReplacedByProductID      string `json:"replacedByProductId"`
		ReplacedByProductURLText string `json:"replacedByProductUrlText"`
	} `json:"redirect"`
	LangUrls           any `json:"langUrls"`
	Pac                any `json:"pac"`
	FulfillmentSummary []struct {
		Fulfillment  string    `json:"fulfillment"`
		StoreID      string    `json:"storeId"`
		DeliveryDate time.Time `json:"deliveryDate"`
		SLA          struct {
			UnitOfMeasure    string `json:"unitOfMeasure"`
			MeasurementValue int    `json:"measurementValue"`
		} `json:"sla"`
		CalculatedSLADays  int      `json:"calculatedSlaDays"`
		FulfillmentMethods []string `json:"fulfillmentMethods"`
		FulfillmentBadge   any      `json:"fulfillmentBadge"`
		IsFreeForWPlus     any      `json:"isFreeForWPlus"`
	} `json:"fulfillmentSummary"`
	ShowExploreOtherConditionsCTA bool   `json:"showExploreOtherConditionsCTA"`
	IsPreowned                    bool   `json:"isPreowned"`
	PreownedCondition             string `json:"preownedCondition"`
	NewConditionProductID         any    `json:"newConditionProductId"`
	DisplayOfferLevelImage        bool   `json:"displayOfferLevelImage"`
	AvailabilityInNearbyStore     any    `json:"availabilityInNearbyStore"`
	SelectedVariantIds            []any  `json:"selectedVariantIds"`
	VariantProductIDMap           struct {
	} `json:"variantProductIdMap"`
	ImageMap struct {
	} `json:"imageMap"`
	VariantsMap struct {
	} `json:"variantsMap"`
}
