package api_response

import (
	"time"
)

type CompanySearch struct {
	TransactionDetail TransactionDetail `json:"transactionDetail"`
	InquiryDetail     InquiryDetail     `json:"inquiryDetail"`
	Candidates        []Candidate       `json:"searchCandidates"`
	Navigators        Navigators        `json:"navigators"`
}

type TransactionDetail struct {
	TransactionID        string    `json:"transactionID"`
	TransactionTimestamp time.Time `json:"transactionTimestamp"`
	InLanguage           string    `json:"inLanguage"`
	ServiceVersion       string    `json:"serviceVersion"`
}

type InquiryDetail struct {
	IsExporter                           bool            `json:"isExporter"`
	TelephoneNumber                      string          `json:"telephoneNumber"`
	PageNumber                           int             `json:"pageNumber"`
	YearlyRevenue                        Revenue         `json:"yearlyRevenue"`
	PostalCode                           string          `json:"postalCode"`
	DunsList                             []string        `json:"dunsList"`
	PageSize                             int             `json:"pageSize"`
	IndustryCodes                        []IndustryCode  `json:"industryCodes"`
	CountryISOCode                       string          `json:"countryISOAlpha2Code"`
	SearchTerm                           string          `json:"searchTerm"`
	ReturnNavigators                     bool            `json:"returnNavigators"`
	USSicV4                              []string        `json:"usSicV4"`
	IsOutOfBusiness                      bool            `json:"isOutOfBusiness"`
	IsImporter                           bool            `json:"isImporter"`
	IsStandalone                         bool            `json:"isStandalone"`
	IsTelephoneDisconnected              bool            `json:"isTelephoneDisconnected"`
	TradeStyleName                       string          `json:"tradeStyleName"`
	FamilyTreeRoles                      []int           `json:"familytreeRolesPlayed"`
	AddressLocality                      string          `json:"addressLocality"`
	AddressRegion                        string          `json:"addressRegion"`
	IsMailUndeliverable                  bool            `json:"isMailUndeliverable"`
	IsMarketable                         bool            `json:"isMarketable"`
	BusinessEntityType                   []int           `json:"businessEntityType"`
	PrimaryName                          string          `json:"primaryName"`
	RegistrationNumbers                  []string        `json:"registrationNumbers"`
	TickerSymbol                         string          `json:"tickerSymbol"`
	Sort                                 []SortCriteria  `json:"sort"`
	StreetAddress                        string          `json:"streetAddressLine1"`
	NumberOfEmployees                    EmployeeNumbers `json:"numberOfEmployees"`
	LocationRadius                       LocationRadius  `json:"locationRadius"`
	GlobalUltimateFamilyTreeMembersCount Range           `json:"globalUltimateFamilyTreeMembersCount"`
	Domain                               string          `json:"domain"`
}

type Revenue struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type IndustryCode struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type SortCriteria struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type EmployeeNumbers struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type LocationRadius struct {
	Radius float64 `json:"radius"`
	Unit   string  `json:"unit"`
}

type Range struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
}

type Candidate struct {
	DisplaySequence int          `json:"displaySequence"`
	Organization    Organization `json:"organization"`
}

type Organization struct {
	Duns                 string             `json:"duns"`
	DunsControlStatus    DunsControlStatus  `json:"dunsControlStatus"`
	PrimaryName          string             `json:"primaryName"`
	PrimaryAddress       PrimaryAddress     `json:"primaryAddress"`
	PrimaryIndustryCodes []IndustryCode     `json:"primaryIndustryCodes"`
	CorporateLinkage     CorporateLinkage   `json:"corporateLinkage"`
	Financials           []Financial        `json:"financials"`
	NumberOfEmployees    []EmployeeNumbers  `json:"numberOfEmployees"`
	IndustryCodes        []IndustryCode     `json:"industryCodes"`
	Telephone            []Telephone        `json:"telephone"`
	BusinessEntityType   BusinessEntityType `json:"businessEntityType"`
	IsStandalone         bool               `json:"isStandalone"`
}

type DunsControlStatus struct {
	OperatingStatus string `json:"operatingStatus"`
	IsOutOfBusiness bool   `json:"isOutOfBusiness"`
}

type PrimaryAddress struct {
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	PostalCode    string `json:"postalCode"`
}

type CorporateLinkage struct {
	GlobalUltimateDuns string `json:"globalUltimateDuns"`
	ParentDuns         string `json:"parentDuns"`
}

type Financial struct {
	Year     int     `json:"year"`
	Revenue  Revenue `json:"revenue"`
	Currency string  `json:"currency"`
}

type Telephone struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type BusinessEntityType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Navigators struct {
	YearlyRevenue      []Navigator `json:"yearlyRevenue"`
	NumberOfEmployees  []Navigator `json:"numberOfEmployees"`
	Industry           []Navigator `json:"industry"`
	BusinessEntityType []Navigator `json:"businessEntityType"`
	FamilyTreeRole     []Navigator `json:"familyTreeRole"`
	Location           Location    `json:"location"`
}

type Navigator struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Location struct {
	Country []Navigator `json:"country"`
	State   []Navigator `json:"state"`
	City    []Navigator `json:"city"`
}
