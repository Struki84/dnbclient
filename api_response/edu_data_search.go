package api_response

type EducationalDataSearch struct {
	Base
	InquiryDetail EducationalDataInquiryDetail `json:"inquiryDetail,omitempty"`
	Navigatiors   []InstitutionNavigators      `json:"navigators,omitempty"`
	Links         []InstitutionLinks           `json:"links,omitempty"`
	Institutions  []Institution                `json:"institutions,omitempty"`
}

type EducationalDataInquiryDetail struct {
	Duns                       string `json:"duns,omitempty"`
	MaxResults                 string `json:"maxResults,omitempty"`
	TradeUp                    string `json:"tradeUp,omitempty"`
	CustomerReference          string `json:"customerReference,omitempty"`
	CandidatesMatchedQuantity  int    `json:"candidatesMatchedQuantity,omitempty"`
	CandidatesReturnedQuantity int    `json:"candidatesReturnedQuantity,omitempty"`
}

type InstitutionNavigators struct {
	Navigators           []InstitutionNavigator `json:"navigators,omitempty"`
	FileType             []InstitutionNavigator `json:"fileType,omitempty"`
	SchoolType           []InstitutionNavigator `json:"schoolType,omitempty"`
	CountryISOAlpha2Code []InstitutionNavigator `json:"countryISOAlpha2Code,omitempty"`
	AddressRegion        []InstitutionNavigator `json:"addressRegion,omitempty"`
	AddressCounty        []InstitutionNavigator `json:"addressCounty,omitempty"`
	AddressLocality      []InstitutionNavigator `json:"addressLocality,omitempty"`
}

type InstitutionNavigator struct {
	Query                     string `json:"query,omitempty"`
	Description               string `json:"description,omitempty"`
	CandidatesMatchedQuantity int    `json:"candidatesMatchedQuantity,omitempty"`
}

type Institution struct {
	Duns                string              `json:"duns,omitempty"`
	InstitutionID       int                 `json:"institutionID,omitempty"`
	InstitutionFullName string              `json:"institutionFullName,omitempty"`
	PostalCode          string              `json:"postalCode,omitempty"`
	MailingAddress      InstitutionAddress  `json:"mailingAddress,omitempty"`
	AddressCountry      InstitutionAddress  `json:"addressCountry,omitempty"`
	AddressRegion       InstitutionAddress  `json:"addressRegion,omitempty"`
	AddressCounty       InstitutionLocality `json:"addressCounty,omitempty"`
	AddressLocality     InstitutionLocality `json:"addressLocality,omitempty"`

	Personnel []struct {
		PersonCompositeID string `json:"personCompositeID,omitempty"`
	} `json:"personnel,omitempty"`
}

type InstitutionAddress struct {
	IsoAlpha2Code   string `json:"isoAlpha2Code,omitempty"`
	AbbreviatedName string `json:"abbreviatedName,omitempty"`
}

type InstitutionLocality struct {
	Name string `json:"name,omitempty"`
}

type InstitutionLinks struct {
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}
