package api_response

// Contact response data
type ContactSearch struct {
	Base
	InquiryDetail ContactInquiryDetail `json:"inquiryDetail,omitempty"`
	Links         CompanyLinks         `json:"links,omitempty"`

	Candidates []struct {
		DisplaySequence int     `json:"displaySequence,omitempty"`
		Contact         Contact `json:"contact,omitempty"`
	} `json:"searchCandidates,omitempty"`
}

type ContactInquiryDetail struct {
	View                                       string   `json:"view,omitempty"`
	SearchTerm                                 string   `json:"searchTerm,omitempty"`
	FamilyTreeScope                            string   `json:"familyTreeScope,omitempty"`
	ContactID                                  string   `json:"contactID,omitempty"`
	ContactEmail                               string   `json:"contactEmail,omitempty"`
	GivenName                                  string   `json:"givenName,omitempty"`
	FamilyName                                 string   `json:"familyName,omitempty"`
	FullName                                   string   `json:"fullName,omitempty"`
	Duns                                       string   `json:"duns,omitempty"`
	CountryISOAlpha2Code                       string   `json:"countryISOAlpha2Code,omitempty"`
	AddressRegion                              string   `json:"addressRegion,omitempty"`
	AddressLocality                            string   `json:"addressLocality,omitempty"`
	PostalCode                                 string   `json:"postalCode,omitempty"`
	TickerSymbol                               string   `json:"tickerSymbol,omitempty"`
	JobTitles                                  []string `json:"jobTitles,omitempty"`
	UsSicV4                                    []string `json:"usSicV4,omitempty"`
	PrimaryName                                string   `json:"primaryName,omitempty"`
	MrcCode                                    []string `json:"mrcCode,omitempty"`
	PageSize                                   int      `json:"pageSize,omitempty"`
	PageNumber                                 int      `json:"pageNumber,omitempty"`
	ReturnNavigators                           bool     `json:"returnNavigators,omitempty"`
	ConfidenceLowerLevelThresholdValue         int      `json:"confidenceLowerLevelThresholdValue,omitempty"`
	CustomerReference                          string   `json:"customerReference,omitempty"`
	TelephoneAccuracyScoreThresholdValue       int      `json:"telephoneAccuracyScoreThresholdValue,omitempty"`
	EmailAccuracyScoreThresholdValue           int      `json:"emailAccuracyScoreThresholdValue,omitempty"`
	HasDirectDial                              bool     `json:"hasDirectDial,omitempty"`
	ReturnManagementResponsibilitiesNavigators bool     `json:"returnManagementResponsibilitiesNavigators,omitempty"`
	ReturnIndustryNavigators                   bool     `json:"returnIndustryNavigators,omitempty"`
	ReturnLocationNavigators                   bool     `json:"returnLocationNavigators,omitempty"`
	LocationNavigatorType                      string   `json:"locationNavigatorType,omitempty"`
	MaxNavigatorBuckets                        int      `json:"maxNavigatorBuckets,omitempty"`
	IncludeSearchResults                       bool     `json:"includeSearchResults,omitempty"`
	CandidatesMatchedQuantity                  int      `json:"candidatesMatchedQuantity,omitempty"`
	CandidatesReturnedQuantity                 int      `json:"candidatesReturnedQuantity,omitempty"`

	Sort []struct {
		Item      string `json:"item,omitempty"`
		Direction string `json:"direction,omitempty"`
	} `json:"sort,omitempty"`

	IndustryCodes []struct {
		Code            []string `json:"code,omitempty"`
		Description     []string `json:"description,omitempty"`
		TypeDescription string   `json:"typeDescription,omitempty"`
		TypeDnbCode     int      `json:"typeDnbCode,omitempty"`
	} `json:"industryCodes,omitempty"`
}

type CompanyLinks struct {
	Last  string `json:"last,omitempty"`
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
}

type Contact struct {
	ID                 string `json:"id,omitempty"`
	GlobalContactKey   string `json:"globalContactKey,omitempty"`
	Email              string `json:"email,omitempty"`
	EmailDomainName    string `json:"emailDomainName,omitempty"`
	GivenName          string `json:"givenName,omitempty"`
	MiddleName         string `json:"middleName,omitempty"`
	FamilyName         string `json:"familyName,omitempty"`
	NamePrefix         string `json:"namePrefix,omitempty"`
	NameSuffix         string `json:"nameSuffix,omitempty"`
	AdditionalName     string `json:"additionalName,omitempty"`
	IsTitleMatched     bool   `json:"isTitleMatched,omitempty"`
	IsSocialVerified   bool   `json:"isSocialVerified,omitempty"`
	DataFreshnessScore int    `json:"dataFreshnessScore,omitempty"`
	ConfidenceLevel    string `json:"confidenceLevel,omitempty"`
	VerifiedDate       string `json:"verifiedDate,omitempty"`

	MatchQualityInformation struct {
		ConfidenceCode   int    `json:"confidenceCode,omitempty"`
		MatchGrade       string `json:"matchGrade,omitempty"`
		MatchDataProfile int    `json:"matchDataProfile,omitempty"`
	} `json:"matchQualityInformation,omitempty"`

	EmailAccuracy struct {
		DeliverabilityScore int    `json:"deliverabilityScore,omitempty"`
		VerifiedDate        string `json:"verifiedDate,omitempty"`
	} `json:"emailAccuracy,omitempty"`

	TitleAccuracy struct {
		AccuracyScore int `json:"accuracyScore,omitempty"`
	} `json:"titleAccuracy,omitempty"`

	Organization struct {
		DUNS        string `json:"duns,omitempty"`
		PrimaryName string `json:"primaryName,omitempty"`
	} `json:"organization,omitempty"`

	ManagementResponsibilities []struct {
		MrcCode string `json:"mrcCode,omitempty"`
	} `json:"managementResponsibilities,omitempty"`

	IndustryCodes []struct {
		Code        []string `json:"code,omitempty"`
		Description []string `json:"description,omitempty"`
		TypeDnbCode int      `json:"typeDnbCode,omitempty"`
	} `json:"industryCodes,omitempty"`

	Telephone []struct {
		TelephoneNumber   string `json:"telephoneNumber,omitempty"`
		TelephoneAccuracy struct {
			AccuracyScore int `json:"accuracyScore,omitempty"`
		} `json:"telephoneAccuracy,omitempty"`
	} `json:"telephone,omitempty"`

	SocialMedia []struct {
		Platform struct {
			Description string `json:"description,omitempty"`
			DnbCode     int    `json:"dnbCode,omitempty"`
		} `json:"platform,omitempty"`
		URL string `json:"url,omitempty"`
	} `json:"socialMedia,omitempty"`

	VanityTitles []struct {
		Title string `json:"title,omitempty"`
	} `json:"vanityTitles,omitempty"`

	JobTitles []struct {
		Title string `json:"title,omitempty"`
	} `json:"jobTitles,omitempty"`
}
