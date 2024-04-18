package dnbclient

type RequestBody struct {
	CompanySearch *CompanySearchRequest
	ContactSearch *ContactSearchRequest
}

// Company Search Request
type CompanySearchRequest struct {
	DUNS                    string   `json:"duns,omitempty"`
	DUNSList                []string `json:"dunsList,omitempty"`
	IsMarketable            bool     `json:"isMarketable,omitempty"`
	IsOutOffBusiness        bool     `json:"isOutOffBusiness,omitempty"`
	IsTelephoneDisconnected bool     `json:"isTelephoneDisconnected,omitempty"`
	IsMailUndeliverable     bool     `json:"isMailUndeliverable,omitempty"`
	SearchTerm              string   `json:"searchTerm,omitempty"`
	PrimaryName             string   `json:"primaryName,omitempty"`
	TradeStyleName          string   `json:"tradeStyleName,omitempty"`
	CountryISOAlpha2Code    string   `json:"countryISOAlpha2Code,omitempty"`
	AddressRegion           string   `json:"addressRegion,omitempty"`
	AddressLocality         string   `json:"addressLocality,omitempty"`
	StreetAddressLine1      string   `json:"streetAddressLine1,omitempty"`
	PostalCode              string   `json:"postalCode,omitempty"`
	TelephoneNumber         string   `json:"telephoneNumber,omitempty"`
	Domain                  string   `json:"domain,omitempty"`
	TickerSymbol            string   `json:"tickerSymbol,omitempty"`
	IsStandalone            bool     `json:"isStandalone,omitempty"`
	IsImporter              bool     `json:"isImporter,omitempty"`
	IsExporter              bool     `json:"isExporter,omitempty"`
	PageNumber              int      `json:"pageNumber,omitempty"`
	PageSize                int      `json:"pageSize,omitempty"`
	ReturnNavigators        bool     `json:"returnNavigators,omitempty"`
	RegistrationNumbers     []string `json:"registrationNumbers,omitempty"`
	BusinessEntityType      []string `json:"businessEntityType,omitempty"`
	FamilytreeRolesPlayed   []string `json:"familytreeRolesPlayed,omitempty"`
	UsSicv4                 []string `json:"usSicv4,omitempty"`
	LocationRadius          struct {
		Lat    float64 `json:"lat"`
		Lng    float64 `json:"lon"`
		Radius float64 `json:"radius"`
		Unit   string  `json:"unit"`
	} `json:"locationRadius,omitempty"`

	NumberOfEmployees struct {
		InformationScope int `json:"informationScope,omitempty"`
		MaximumValue     int `json:"maximumValue,omitempty"`
		MinimumValue     int `json:"minimumValue,omitempty"`
	} `json:"numberOfEmployees,omitempty"`

	YearlyRevenue struct {
		MaximumValue int `json:"maximumValue,omitempty"`
		MinimumValue int `json:"minimumValue,omitempty"`
	}

	GlobalUltimateFamilyTreeMembersCount struct {
		MaximumValue int `json:"maximumValue,omitempty"`
		MinimumValue int `json:"minimumValue,omitempty"`
	} `json:"globalUltimateFamilyTreeMembersCount,omitempty"`

	IndustryCodes []struct {
		TypeDnbCode string   `json:"typeDnbCode,omitempty"`
		Description string   `json:"description,omitempty"`
		Code        []string `json:"code,omitempty"`
	} `json:"industryCodes,omitempty"`
}

// Typehead Search Request
type TypeheadSearchRequest struct {
	SearchTerm               string  `json:"searchTerm"`
	CountryISOAlpha2Code     string  `json:"countryISOAlpha2Code,omitempty"`
	IsOutOfBusiness          bool    `json:"isOutOfBusiness,omitempty"`
	IsMarketable             bool    `json:"isMarketable,omitempty"`
	IsDelisted               bool    `json:"isDelisted"`
	IsMailUndeliverable      bool    `json:"isMailUndeliverable,omitempty"`
	AddressLocality          string  `json:"addressLocality,omitempty"`
	AddressRegion            string  `json:"addressRegion,omitempty"`
	StreetAddressLine1       string  `json:"streetAddressLine1,omitempty"`
	PostalCode               string  `json:"postalCode,omitempty"`
	RadiusLat                float64 `json:"radiusLat,omitempty"`
	RadiusLon                float64 `json:"radiusLon,omitempty"`
	RadiusPostalCode         string  `json:"radiusPostalCode,omitempty"`
	RadiusDistance           float64 `json:"radiusDistance,omitempty"`
	RadiusUnit               string  `json:"radiusUnit,omitempty"`
	CandidateMaximumQuantity int     `json:"candidateMaximumQuantity,omitempty"`
	CustomerReference        string  `json:"customerReference,omitempty"`
}

// Contact Search Request
type ContactSearchRequest struct {
	ContactID            string         `json:"contactID,omitempty"`
	ContactEmail         string         `json:"contactEmail,omitempty"`
	GivenName            string         `json:"givenName,omitempty"`
	FamilyName           string         `json:"familyName,omitempty"`
	JobTitles            []string       `json:"jobTitles,omitempty"`
	Duns                 string         `json:"duns,omitempty"`
	PrimaryName          string         `json:"primaryName,omitempty"`
	AddressLocality      string         `json:"addressLocality,omitempty"`
	AddressRegion        string         `json:"addressRegion,omitempty"`
	PostalCode           string         `json:"postalCode,omitempty"`
	CountryISOAlpha2Code string         `json:"countryISOAlpha2Code,omitempty"`
	UsSicV4              []string       `json:"usSicV4,omitempty"`
	MrcCode              []string       `json:"mrcCode,omitempty"`
	View                 string         `json:"view,omitempty"`
	HasDirectDial        bool           `json:"hasDirectDial,omitempty"`
	IndustryCodes        []IndustryCode `json:"industryCodes,omitempty"`
	ReturnNavigators     bool           `json:"returnNavigators,omitempty"`
	PageNumber           int            `json:"pageNumber,omitempty"`
	PageSize             int            `json:"pageSize,omitempty"`
	Sort                 []SortItem     `json:"sort,omitempty"`
}

type IndustryCode struct {
	TypeDnbCode int      `json:"typeDnbCode,omitempty"`
	Description []string `json:"description,omitempty"`
	Code        []string `json:"code,omitempty"`
}

type SortItem struct {
	Item      string `json:"item,omitempty"`
	Direction string `json:"direction,omitempty"`
}
