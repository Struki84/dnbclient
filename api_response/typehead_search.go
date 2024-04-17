// package api_response
//
// type Response struct {
// 	TransactionDetail          TransactionDetail `json:"transactionDetail,omitempty"`
// 	InquiryDetail              InquiryDetail     `json:"inquiryDetail,omitempty"`
// 	CandidatesReturnedQuantity int               `json:"candidatesReturnedQuantity,omitempty"`
// 	CandidatesMatchedQuantity  int               `json:"candidatesMatchedQuantity,omitempty"`
// 	SearchCandidates           []Candidate       `json:"searchCandidates,omitempty"`
// }
//
//
// type InquiryDetail struct {
// 	CountryISOAlpha2Code string `json:"countryISOAlpha2Code,omitempty"`
// 	SearchTerm           string `json:"searchTerm,omitempty"`
// }
//
// type Candidate struct {
// 	DisplaySequence int          `json:"displaySequence,omitempty"`
// 	Organization    Organization `json:"organization,omitempty"`
// }
//
// type Organization struct {
// 	Duns                 string                `json:"duns,omitempty"`
// 	DunsControlStatus    DunsControlStatus     `json:"dunsControlStatus,omitempty"`
// 	PrimaryName          string                `json:"primaryName,omitempty"`
// 	PrimaryAddress       PrimaryAddress        `json:"primaryAddress,omitempty"`
// 	CorporateLinkage     CorporateLinkage      `json:"corporateLinkage,omitempty"`
// 	Financials           []Financial           `json:"financials,omitempty"`
// 	TradeStyleNames      []TradeStyleName      `json:"tradeStyleNames,omitempty"`
// 	PrimaryIndustryCodes []PrimaryIndustryCode `json:"primaryIndustryCodes,omitempty"`
// }
//
// type DunsControlStatus struct {
// 	IsOutOfBusiness bool `json:"isOutOfBusiness,omitempty"`
// }
//
// type PrimaryAddress struct {
// 	AddressCountry  AddressCountry  `json:"addressCountry,omitempty"`
// 	StreetAddress   StreetAddress   `json:"streetAddress,omitempty"`
// 	AddressLocality AddressLocality `json:"addressLocality,omitempty"`
// 	AddressRegion   AddressRegion   `json:"addressRegion,omitempty"`
// }
//
// type AddressCountry struct {
// 	IsoAlpha2Code string `json:"isoAlpha2Code,omitempty"`
// }
//
// type StreetAddress struct {
// 	Line1 string `json:"line1,omitempty"`
// }
//
// type AddressLocality struct {
// 	Name string `json:"name,omitempty"`
// }
//
// type AddressRegion struct {
// 	Name string `json:"name,omitempty"`
// }
//
// type CorporateLinkage struct {
// 	IsBranch bool `json:"isBranch,omitempty"`
// }
//
// type Financial struct {
// 	YearlyRevenue []YearlyRevenue `json:"yearlyRevenue,omitempty"`
// }
//
// type YearlyRevenue struct {
// 	Value    float64 `json:"value,omitempty"`
// 	Currency string  `json:"currency,omitempty"`
// }
//
// type TradeStyleName struct {
// 	Name     string `json:"name,omitempty"`
// 	Priority int    `json:"priority,omitempty"`
// }
//
// type PrimaryIndustryCode struct {
// 	UsSicV4            string `json:"usSicV4,omitempty"`
// 	UsSicV4Description string `json:"usSicV4Description,omitempty"`
// }
