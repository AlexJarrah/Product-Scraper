package nike

type NikeData struct {
	Props        Props  `json:"props"`
	Page         string `json:"page"`
	Query        Query  `json:"query"`
	BuildID      string `json:"buildId"`
	AssetPrefix  string `json:"assetPrefix"`
	IsFallback   bool   `json:"isFallback"`
	Gssp         bool   `json:"gssp"`
	CustomServer bool   `json:"customServer"`
}
type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}
type Meta struct {
	Content            string `json:"content"`
	DataAPIBranchName  string `json:"data-api-branch-name,omitempty"`
	DataAPIBuildNumber string `json:"data-api-build-number,omitempty"`
	DataAPICommitSha   string `json:"data-api-commit-sha,omitempty"`
	Name               string `json:"name,omitempty"`
	Property           string `json:"property,omitempty"`
}
type MetaTags struct {
	Link    []Link `json:"link"`
	Meta    []Meta `json:"meta"`
	Title   string `json:"title"`
	URLPath string `json:"urlPath"`
}
type Request struct {
}
type Global struct {
	MetaTags MetaTags `json:"metaTags"`
	Request  Request  `json:"request"`
}

type CountryNames struct {
	En string `json:"en"`
}

type Banner struct {
	Rendered      bool  `json:"rendered"`
	RequestFailed bool  `json:"requestFailed"`
	Results       []any `json:"results"`
}
type User struct {
	Initialized bool `json:"initialized"`
	IsLoggedIn  bool `json:"isLoggedIn"`
	IsSwoosh    bool `json:"isSwoosh"`
}
type APIMetaData struct {
	XBranchName    string `json:"x-branch-name"`
	XBuildNumber   string `json:"x-build-number"`
	XCommitSha     string `json:"x-commit-sha"`
	XSearchOptions any    `json:"x-search-options"`
}
type Analytics struct {
	LegacyPageName string `json:"legacyPageName"`
}
type ClearAllFilters struct {
	Crawlable              any `json:"crawlable"`
	Navigation             any `json:"navigation"`
	NavigationAttributeIds any `json:"navigationAttributeIds"`
}
type Navigation struct {
	CanonicalURL string `json:"canonicalUrl"`
	PageURL      string `json:"pageUrl"`
	Path         string `json:"path"`
}
type Options struct {
	AlternateName          string     `json:"alternateName"`
	AttributeID            string     `json:"attributeId"`
	Crawlable              bool       `json:"crawlable"`
	StyleCat               any        `json:"styleCat"`
	DisplayText            string     `json:"displayText"`
	Navigation             Navigation `json:"navigation"`
	NavigationAttributeIds []string   `json:"navigationAttributeIds"`
	ResultCount            int        `json:"resultCount"`
	SearchTerms            string     `json:"searchTerms"`
	Selected               bool       `json:"selected"`
}
type Filters struct {
	AlternateName string    `json:"alternateName"`
	AttributeID   string    `json:"attributeId"`
	DisplayText   string    `json:"displayText"`
	Expanded      bool      `json:"expanded"`
	Options       []Options `json:"options"`
	ResultCount   int       `json:"resultCount"`
	SelectType    string    `json:"selectType"`
	Selected      bool      `json:"selected"`
	ShowMoreAfter int       `json:"showMoreAfter"`
}
type FacetNav struct {
	Analytics        Analytics       `json:"analytics"`
	SelectedConcepts []any           `json:"selectedConcepts"`
	Categories       []any           `json:"categories"`
	ClearAllFilters  ClearAllFilters `json:"clearAllFilters"`
	Filters          []Filters       `json:"filters"`
	Breadcrumbs      []any           `json:"breadcrumbs"`
}
type SearchSummary struct {
	OriginalTerms string `json:"originalTerms"`
}
type PageData struct {
	Prev           string        `json:"prev"`
	Next           string        `json:"next"`
	TotalPages     int           `json:"totalPages"`
	TotalResources int           `json:"totalResources"`
	SearchSummary  SearchSummary `json:"searchSummary"`
}
type Images struct {
	PortraitURL string `json:"portraitURL"`
	SquarishURL string `json:"squarishURL"`
}
type Price struct {
	Currency               string `json:"currency"`
	CurrentPrice           int    `json:"currentPrice"`
	Discounted             bool   `json:"discounted"`
	EmployeePrice          int    `json:"employeePrice"`
	FullPrice              int    `json:"fullPrice"`
	MinimumAdvertisedPrice any    `json:"minimumAdvertisedPrice"`
}
type Colorways struct {
	ColorDescription  string `json:"colorDescription"`
	Images            Images `json:"images"`
	PdpURL            string `json:"pdpUrl"`
	Price             Price  `json:"price"`
	CloudProductID    string `json:"cloudProductId"`
	InStock           bool   `json:"inStock"`
	IsComingSoon      bool   `json:"isComingSoon"`
	IsBestSeller      bool   `json:"isBestSeller"`
	IsExcluded        bool   `json:"isExcluded"`
	IsJPStrikethrough any    `json:"isJPStrikethrough"`
	IsLaunch          bool   `json:"isLaunch"`
	IsMemberExclusive bool   `json:"isMemberExclusive"`
	IsNew             bool   `json:"isNew"`
	Label             string `json:"label"`
	Pid               string `json:"pid"`
	PrebuildID        any    `json:"prebuildId"`
	ProductInstanceID any    `json:"productInstanceId"`
}
type Products struct {
	CardType          string      `json:"cardType"`
	CloudProductID    string      `json:"cloudProductId"`
	ColorDescription  string      `json:"colorDescription"`
	Colorways         []Colorways `json:"colorways"`
	Customizable      bool        `json:"customizable"`
	HasExtendedSizing bool        `json:"hasExtendedSizing"`
	ID                string      `json:"id"`
	Images            Images      `json:"images"`
	InStock           bool        `json:"inStock"`
	IsComingSoon      bool        `json:"isComingSoon"`
	IsBestSeller      bool        `json:"isBestSeller"`
	IsExcluded        bool        `json:"isExcluded"`
	IsGiftCard        bool        `json:"isGiftCard"`
	IsJersey          bool        `json:"isJersey"`
	IsJPStrikethrough any         `json:"isJPStrikethrough"`
	IsLaunch          bool        `json:"isLaunch"`
	IsMemberExclusive bool        `json:"isMemberExclusive"`
	IsNBA             bool        `json:"isNBA"`
	IsNFL             bool        `json:"isNFL"`
	IsSustainable     bool        `json:"isSustainable"`
	Label             string      `json:"label"`
	NbyColorway       any         `json:"nbyColorway"`
	Pid               string      `json:"pid"`
	PrebuildID        any         `json:"prebuildId"`
	Price             Price       `json:"price"`
	ProductInstanceID any         `json:"productInstanceId"`
	ProductType       string      `json:"productType"`
	Properties        any         `json:"properties"`
	SalesChannel      []string    `json:"salesChannel"`
	Subtitle          string      `json:"subtitle"`
	Title             string      `json:"title"`
	URL               string      `json:"url"`
}
type Attributes struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}
type Audiences struct {
	Name       string `json:"name"`
	Conditions string `json:"conditions"`
	ID         string `json:"id"`
}
type Events struct {
	ID            string   `json:"id"`
	ExperimentIds []string `json:"experimentIds"`
	Key           string   `json:"key"`
}
type TypedAudiences struct {
	Name       string `json:"name"`
	Conditions []any  `json:"conditions"`
	ID         string `json:"id"`
}
type Variables struct {
	ID           string `json:"id"`
	Key          string `json:"key"`
	Type         string `json:"type"`
	DefaultValue string `json:"defaultValue"`
}
type FeatureFlags struct {
	ID            string      `json:"id"`
	Key           string      `json:"key"`
	ExperimentIds []any       `json:"experimentIds"`
	RolloutID     string      `json:"rolloutId"`
	Variables     []Variables `json:"variables"`
}
type ForcedVariations struct {
}
type Variations struct {
	ID             string `json:"id"`
	Key            string `json:"key"`
	Variables      []any  `json:"variables"`
	FeatureEnabled bool   `json:"featureEnabled"`
}
type Experiments struct {
	ForcedVariations   ForcedVariations `json:"forcedVariations"`
	ID                 string           `json:"id"`
	Key                string           `json:"key"`
	Status             string           `json:"status"`
	TrafficAllocation  []any            `json:"trafficAllocation"`
	Variations         []Variations     `json:"variations"`
	LayerID            string           `json:"layerId"`
	AudienceIds        []string         `json:"audienceIds"`
	AudienceConditions []string         `json:"audienceConditions"`
}
type Rollouts struct {
	ID          string        `json:"id"`
	Experiments []Experiments `json:"experiments"`
}
type TrafficAllocation struct {
	EntityID   string `json:"entityId"`
	EndOfRange int    `json:"endOfRange"`
}
type Datafile struct {
	AccountID      string           `json:"accountId"`
	ProjectID      string           `json:"projectId"`
	Revision       string           `json:"revision"`
	Attributes     []Attributes     `json:"attributes"`
	Audiences      []Audiences      `json:"audiences"`
	Version        string           `json:"version"`
	Events         []Events         `json:"events"`
	Integrations   []any            `json:"integrations"`
	AnonymizeIP    bool             `json:"anonymizeIP"`
	BotFiltering   bool             `json:"botFiltering"`
	TypedAudiences []TypedAudiences `json:"typedAudiences"`
	Variables      []any            `json:"variables"`
	EnvironmentKey string           `json:"environmentKey"`
	SdkKey         string           `json:"sdkKey"`
	FeatureFlags   []FeatureFlags   `json:"featureFlags"`
	Rollouts       []Rollouts       `json:"rollouts"`
	Experiments    []Experiments    `json:"experiments"`
	Groups         []any            `json:"groups"`
}
type Wall struct {
	APIMetaData            APIMetaData `json:"apiMetaData"`
	CanonicalURL           string      `json:"canonicalUrl"`
	CmsContent             any         `json:"cmsContent"`
	FacetNav               FacetNav    `json:"facetNav"`
	IsOptimizelyCompleted  bool        `json:"isOptimizelyCompleted"`
	Loading                bool        `json:"loading"`
	PageData               PageData    `json:"pageData"`
	Products               []Products  `json:"products"`
	RelatedCategories      any         `json:"relatedCategories"`
	RelatedContent         any         `json:"relatedContent"`
	SearchTerm             string      `json:"searchTerm"`
	SelectedFiltersCount   int         `json:"selectedFiltersCount"`
	SortBy                 string      `json:"sortBy"`
	Title                  any         `json:"title"`
	TraceID                string      `json:"traceId"`
	Uuids                  any         `json:"uuids"`
	VisibleFilterNavGroups []any       `json:"visibleFilterNavGroups"`
	Datafile               Datafile    `json:"datafile"`
	PrimaryHeading         any         `json:"primaryHeading"`
}
type Flags struct {
	AbTests                 bool `json:"abTests"`
	CookieSettings          bool `json:"cookieSettings"`
	Debug                   bool `json:"debug"`
	EnableMAP               bool `json:"enableMAP"`
	EnablePercentOffDisplay bool `json:"enablePercentOffDisplay"`
	ForcePriceDecimals      bool `json:"forcePriceDecimals"`
	GlobalNavUseGeoPrivacy  bool `json:"globalNavUseGeoPrivacy"`
	PulseInsights           bool `json:"pulseInsights"`
	Recommendations         bool `json:"recommendations"`
	SendPIDsOnWall          bool `json:"sendPIDsOnWall"`
	ShowGeoMismatch         bool `json:"showGeoMismatch"`
	ShowFindInStore         bool `json:"showFindInStore"`
	ShowJPStrikethrough     bool `json:"showJPStrikethrough"`
	ShowPromoBanner         bool `json:"showPromoBanner"`
	ShowRelatedCategories   bool `json:"showRelatedCategories"`
	ShowRelatedContent      bool `json:"showRelatedContent"`
	ShowSEOCopy             bool `json:"showSEOCopy"`
	ShowUniteTimers         bool `json:"showUniteTimers"`
	SwooshEligibleGeo       bool `json:"swooshEligibleGeo"`
	UseDotcomNav            bool `json:"useDotcomNav"`
	UseOptimizelyX          bool `json:"useOptimizelyX"`
	UsePlainSaleMessaging   bool `json:"usePlainSaleMessaging"`
	UseServiceCanonicalURL  bool `json:"useServiceCanonicalUrl"`
	UseUniteStaging         bool `json:"useUniteStaging"`
	EnableNikeShopModals    bool `json:"enableNikeShopModals"`
	WallFavorites           bool `json:"wallFavorites"`
	WallImpressions         bool `json:"wallImpressions"`
	WallNoResultsCarousel   bool `json:"wallNoResultsCarousel"`
	XSearchOptionsHeader    bool `json:"xSearchOptionsHeader"`
}
type CacheDuration struct {
	Num404  int `json:"404"`
	Default int `json:"default"`
	Perf    int `json:"perf"`
	Wall    int `json:"wall"`
}
type Config struct {
	ChannelID     string        `json:"channelId"`
	Division      string        `json:"division"`
	Domain        string        `json:"domain"`
	FacebookAppID string        `json:"facebookAppId"`
	Name          string        `json:"name"`
	Platform      string        `json:"platform"`
	Version       string        `json:"version"`
	Dark          bool          `json:"dark"`
	Env           string        `json:"env"`
	Preview       bool          `json:"preview"`
	GqlHash       string        `json:"gqlHash"`
	Flags         Flags         `json:"flags"`
	CacheDuration CacheDuration `json:"cacheDuration"`
}
type Console struct {
	Providers []any `json:"providers"`
}
type Privacy struct {
	ArePerformanceCookiesAllowed bool `json:"arePerformanceCookiesAllowed"`
	AreSocialCookiesAllowed      bool `json:"areSocialCookiesAllowed"`
	HasEUCookie                  bool `json:"hasEUCookie"`
}
type Device struct {
	DeviceWidth string `json:"deviceWidth"`
}
type Common struct {
	Channel                  string `json:"channel"`
	GlobalNavFixed           bool   `json:"globalNavFixed"`
	GlobalNavHidden          bool   `json:"globalNavHidden"`
	GlobalNavPeekabooEnabled bool   `json:"globalNavPeekabooEnabled"`
	GlobalNavPeekabooVisible bool   `json:"globalNavPeekabooVisible"`
}
type StoreLocations struct {
	FilteredByStore bool   `json:"filteredByStore"`
	PostalCode      string `json:"postalCode"`
	StoreList       []any  `json:"storeList"`
}
type InitialState struct {
	Global         Global         `json:"global"`
	Banner         Banner         `json:"banner"`
	User           User           `json:"User"`
	Wall           Wall           `json:"Wall"`
	Config         Config         `json:"Config"`
	Console        Console        `json:"Console"`
	Experiments    Experiments    `json:"Experiments"`
	Privacy        Privacy        `json:"Privacy"`
	Device         Device         `json:"Device"`
	Common         Common         `json:"Common"`
	StoreLocations StoreLocations `json:"StoreLocations"`
}
type PageProps struct {
	InitialState     InitialState `json:"initialState"`
	CloudURLFragment string       `json:"cloudUrlFragment"`
}
type Props struct {
	PageProps PageProps `json:"pageProps"`
	NSsp      bool      `json:"__N_SSP"`
}
type Query struct {
	Q   string `json:"q"`
	Vst string `json:"vst"`
}
