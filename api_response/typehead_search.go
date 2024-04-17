package api_response

type TypeheadSearch struct {
	Base
	InquiryDetail TypeheadInquiryDetail `json:"inquiryDetail,omitempty"`

	SearchCandidates []struct {
		DisplaySequence int                  `json:"displaySequence,omitempty"`
		Organization    TypeheadOrganization `json:"organization,omitempty"`
	} `json:"searchCandidates,omitempty"`
}

type TypeheadInquiryDetail struct {
	SearchTerm                 string  `json:"searchTerm,omitempty"`
	CountryISOAlpha2Code       string  `json:"countryISOAlpha2Code,omitempty"`
	IsOutOfBusiness            bool    `json:"isOutOfBusiness,omitempty"`
	IsMarketable               bool    `json:"isMarketable,omitempty"`
	IsDelisted                 bool    `json:"isDelisted,omitempty"`
	IsMailUndeliverable        bool    `json:"isMailUndeliverable,omitempty"`
	AddressLocality            string  `json:"addressLocality,omitempty"`
	AddressRegion              string  `json:"addressRegion,omitempty"`
	StreetAddressLine1         string  `json:"streetAddressLine1,omitempty"`
	PostalCode                 string  `json:"postalCode,omitempty"`
	RadiusLat                  float64 `json:"radiusLat,omitempty"`
	RadiusLon                  float64 `json:"radiusLon,omitempty"`
	RadiusPostalCode           string  `json:"radiusPostalCode,omitempty"`
	RadiusDistance             float64 `json:"radiusDistance,omitempty"`
	RadiusUnit                 string  `json:"radiusUnit,omitempty"`
	CandidateMaximumQuantity   int     `json:"candidateMaximumQuantity,omitempty"`
	CustomerReference          string  `json:"customerReference,omitempty"`
	CandidatesMatchedQuantity  int     `json:"candidatesMatchedQuantity,omitempty"`
	CandidatesReturnedQuantity int     `json:"candidatesReturnedQuantity,omitempty"`
}

type TypeheadOrganization struct {
	Duns        string `json:"duns,omitempty"`
	PrimaryName string `json:"primaryName,omitempty"`

	DunsControlStatus struct {
		IsOutOfBusiness bool `json:"isOutOfBusiness,omitempty"`
	} `json:"dunsControlStatus,omitempty"`

	PrimaryAddress struct {
		AddressCountry struct {
			IsoAlpha2Code string `json:"isoAlpha2Code,omitempty"`
		} `json:"addressCountry,omitempty"`

		StreetAddress struct {
			Line1 string `json:"line1,omitempty"`
		} `json:"streetAddress,omitempty"`

		AddressLocality struct {
			Name string `json:"name,omitempty"`
		} `json:"addressLocality,omitempty"`

		AddressRegion struct {
			Name string `json:"name,omitempty"`
		} `json:"addressRegion,omitempty"`
	} `json:"primaryAddress,omitempty"`

	CorporateLinkage struct {
		IsBranch bool `json:"isBranch,omitempty"`
	} `json:"corporateLinkage,omitempty"`

	Financials []struct {
		YearlyRevenue []struct {
			Value    float64 `json:"value,omitempty"`
			Currency string  `json:"currency,omitempty"`
		} `json:"yearlyRevenue,omitempty"`
	} `json:"financials,omitempty"`

	TradeStyleNames []struct {
		Name     string `json:"name,omitempty"`
		Priority int    `json:"priority,omitempty"`
	} `json:"tradeStyleNames,omitempty"`

	PrimaryIndustryCodes []struct {
		UsSicV4            string `json:"usSicV4,omitempty"`
		UsSicV4Description string `json:"usSicV4Description,omitempty"`
	} `json:"primaryIndustryCodes,omitempty"`
}
