package api_response

type CompanySearch struct {
	Base
	InquiryDetail CompanyInquiryDetail `json:"inquiryDetail,omitempty"`
	Navigators    CompanyNavigators    `json:"navigators,omitempty"`

	Candidates []struct {
		DisplaySequence int          `json:"displaySequence,omitempty"`
		Organization    Organization `json:"organization,omitempty"`
	} `json:"searchCandidates,omitempty"`
}

type CompanyInquiryDetail struct {
	IsExporter              bool     `json:"isExporter,omitempty"`
	TelephoneNumber         string   `json:"telephoneNumber,omitempty"`
	PageNumber              int      `json:"pageNumber,omitempty"`
	PostalCode              string   `json:"postalCode,omitempty"`
	DunsList                []string `json:"dunsList,omitempty"`
	PageSize                int      `json:"pageSize,omitempty"`
	CountryISOCode          string   `json:"countryISOAlpha2Code,omitempty"`
	SearchTerm              string   `json:"searchTerm,omitempty"`
	ReturnNavigators        bool     `json:"returnNavigators,omitempty"`
	USSicV4                 []string `json:"usSicV4,omitempty"`
	IsOutOfBusiness         bool     `json:"isOutOfBusiness,omitempty"`
	IsImporter              bool     `json:"isImporter,omitempty"`
	IsStandalone            bool     `json:"isStandalone,omitempty"`
	IsTelephoneDisconnected bool     `json:"isTelephoneDisconnected,omitempty"`
	TradeStyleName          string   `json:"tradeStyleName,omitempty"`
	FamilyTreeRoles         []int    `json:"familytreeRolesPlayed,omitempty"`
	AddressLocality         string   `json:"addressLocality,omitempty"`
	AddressRegion           string   `json:"addressRegion,omitempty"`
	IsMailUndeliverable     bool     `json:"isMailUndeliverable,omitempty"`
	IsMarketable            bool     `json:"isMarketable,omitempty"`
	BusinessEntityType      []int    `json:"businessEntityType,omitempty"`
	PrimaryName             string   `json:"primaryName,omitempty"`
	RegistrationNumbers     []string `json:"registrationNumbers,omitempty"`
	TickerSymbol            string   `json:"tickerSymbol,omitempty"`
	StreetAddress           string   `json:"streetAddressLine1,omitempty"`
	Domain                  string   `json:"domain,omitempty"`

	IndustryCodes []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
		TypeDnbCode string `json:"typeDnbCode"`
	} `json:"industryCodes"`

	Sort []struct {
		Field     string `json:"field"`
		Direction string `json:"direction"`
	} `json:"sort"`

	NumberOfEmployees struct {
		Value int `json:"value"`
	} `json:"numberOfEmployees"`

	LocationRadius struct {
		Radius float64 `json:"radius"`
		Unit   string  `json:"unit"`
	} `json:"locationRadius"`

	GlobalUltimateFamilyTreeMembersCount struct {
		Minimum int `json:"minimum"`
		Maximum int `json:"maximum"`
	} `json:"globalUltimateFamilyTreeMembersCount"`

	YearlyRevenue struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"yearlyRevenue,omitempty"`
}

type Organization struct {
	Duns         string `json:"duns"`
	IsStandalone bool   `json:"isStandalone"`
	PrimaryName  string `json:"primaryName"`

	DunsControlStatus struct {
		OperatingStatus string `json:"operatingStatus"`
		IsOutOfBusiness bool   `json:"isOutOfBusiness"`
	} `json:"dunsControlStatus"`

	PrimaryAddress struct {
		StreetAddress string `json:"streetAddress"`
		City          string `json:"city"`
		State         string `json:"state"`
		Country       string `json:"country"`
		PostalCode    string `json:"postalCode"`
	} `json:"primaryAddress"`

	CorporateLinkage struct {
		GlobalUltimateDuns string `json:"globalUltimateDuns"`
		ParentDuns         string `json:"parentDuns"`
	} `json:"corporateLinkage"`

	BusinessEntityType struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"businessEntityType"`

	Financials []struct {
		Year          int `json:"year"`
		YearlyRevenue struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"yearlyRevenue"`
	} `json:"financials"`

	PrimaryIndustryCodes []struct {
		Code            []string `json:"code,omitempty"`
		Description     []string `json:"description,omitempty"`
		TypeDescription string   `json:"typeDescription,omitempty"`
		TypeDnbCode     int      `json:"typeDnbCode,omitempty"`
	} `json:"primaryIndustryCodes,omitempty"`

	NumberOfEmployees []struct {
		Value int `json:"value"`
	} `json:"numberOfEmployees"`

	IndustryCodes []struct {
		Code            []string `json:"code,omitempty"`
		Description     []string `json:"description,omitempty"`
		TypeDescription string   `json:"typeDescription,omitempty"`
		TypeDnbCode     int      `json:"typeDnbCode,omitempty"`
	} `json:"industryCodes,omitempty"`

	Telephone []struct {
		Type   string `json:"type"`
		Number string `json:"number"`
	} `json:"telephone"`
}

type CompanyNavigators struct {
	YearlyRevenue      []CompanyNavigator `json:"yearlyRevenue"`
	NumberOfEmployees  []CompanyNavigator `json:"numberOfEmployees"`
	Industry           []CompanyNavigator `json:"industry"`
	BusinessEntityType []CompanyNavigator `json:"businessEntityType"`
	FamilyTreeRole     []CompanyNavigator `json:"familyTreeRole"`
	Location           struct {
		Country []CompanyNavigator `json:"country"`
		State   []CompanyNavigator `json:"state"`
		City    []CompanyNavigator `json:"city"`
	} `json:"location"`
}

type CompanyNavigator struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
