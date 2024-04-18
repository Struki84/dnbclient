package api_response

type CompetitorsSearch struct {
	Base
	InquiryDetail CompetitorsInquiryDetail `json:"inquiryDetail,omitempty"`
	Competitors   []Competitor             `json:"competitors,omitempty"`
}

type CompetitorsInquiryDetail struct {
	Duns                       string `json:"duns,omitempty"`
	MaxResults                 string `json:"maxResults,omitempty"`
	TradeUp                    string `json:"tradeUp,omitempty"`
	CustomerReference          string `json:"customerReference,omitempty"`
	CandidatesMatchedQuantity  int    `json:"candidatesMatchedQuantity,omitempty"`
	CandidatesReturnedQuantity int    `json:"candidatesReturnedQuantity,omitempty"`
}

type Competitor struct {
	ConsolidatedEmployeeCount int     `json:"consolidatedEmployeeCount,omitempty"`
	SalesRevenue              float64 `json:"salesRevenue,omitempty"`
	SalesRevenueCurrency      string  `json:"salesRevenueCurrency,omitempty"`
	IssuedShareCapitalAmount  float64 `json:"issuedShareCapitalAmount,omitempty"`

	CorporateLinkage struct {
		GlobalUltimate struct {
			Duns string `json:"duns,omitempty"`
		} `json:"globalUltimate,omitempty"`
		Parent struct {
			Duns string `json:"duns,omitempty"`
		} `json:"parent,omitempty"`
	} `json:"corporateLinkage,omitempty"`

	PrimaryAddress struct {
		AddressCountry struct {
			IsoAlpha2Code string `json:"isoAlpha2Code,omitempty"`
			Name          string `json:"name,omitempty"`
		} `json:"addressCountry,omitempty"`
		AddressLocality struct {
			Name string `json:"name,omitempty"`
		} `json:"addressLocality,omitempty"`
		AddressRegion struct {
			Name            string `json:"name,omitempty"`
			AbbreviatedName string `json:"abbreviatedName,omitempty"`
		} `json:"addressRegion,omitempty"`
	} `json:"primaryAddress,omitempty"`
}
