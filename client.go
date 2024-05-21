package dnbclient

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/struki84/dnbclient/api_response"
)

const (

	// API ENDPOINTS

	// V1 Base API url
	BaseURLV1 = "https://plus.dnb.com/v1"

	// V2 base API url
	BaseURLV3 = "https://plus.dnb.com/v3"

	//Authorization endpoint for generating API token
	AuthURL = "/token"

	// Crteria compant search endpoint
	CriteriaSearchURL = "/search/criteria"

	// Typehead company search endpoint
	TypeheadSearchURL = "/search/typehead"

	// Company list search endpoint
	CompanyListURL = "/search/companyList"

	// Contact search endpoint
	ContactSearchURL = "/search/contact"
)

var (
	ErrMissingAPIKey        = errors.New("api token is required")
	ErrGetTokenFailed       = errors.New("get token failed")
	ErrSearchCriteriaFailed = errors.New("search criteria failed")
	ErrCompanyListFailed    = errors.New("company list search failed")
	ErrTypeheadSearchFailed = errors.New("typehead search failed")
	ErrContactSearchFailed  = errors.New("contact search failed")
	ErrGetContactsFailed    = errors.New("get contacts failed")
	ErrNoSearchResults      = errors.New("no search results found")
	ErrRequestFailed        = errors.New("http request failed with error")
)

type Client struct {
	username    string
	password    string
	ApiKey      string
	ApiSecret   string
	apiToken    string
	options     []ClientOptions
	BaseURL     string
	RequestBody *RequestBody
}

// NewClient creates a new DNB client
//
// Parameters
// - options: allows ihntial configuration of the client
//
// Returns
// - client: DNB client
// - error: error if any
func NewClient(options ...ClientOptions) (*Client, error) {
	client := &Client{
		BaseURL:     BaseURLV1,
		RequestBody: &RequestBody{},
	}

	return client, nil
}

// GetToken generates B&B Direct+ API token, requires D&B api key and
// api secret to generate token.
//
// # Parameters:
//
// - ctx
//
// - options: variatic parameter that allows passing in client options
// like api key and api secret username and password in the function
//
// # Returns
//
// - token: API token string used in all consequent API requests
//
// - error: error if any
//
// # Documentation
//
// - https://directplus.documentation.dnb.com/openAPI.html?apiID=authenticationV3
func (client *Client) GetToken(ctx context.Context, options ...ClientOptions) (string, error) {
	client.loadOptions(options...)

	credentials := client.ApiKey + ":" + client.ApiSecret
	client.apiToken = base64.StdEncoding.EncodeToString([]byte(credentials))

	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")

	reqBody := formData.Encode()

	// var reqData struct {
	// 	GrantType string `json:"grant_type"`
	// }
	//
	// reqData.GrantType = "client_credentials"
	//
	// reqBytes, err := json.Marshal(reqData)
	// if err != nil {
	// 	return "", fmt.Errorf("%w, %w", ErrGetTokenFailed, err)
	// }

	reqURL := client.BaseURL + AuthURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, strings.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("%w, %w", ErrGetTokenFailed, err)
	}

	req.Header.Add("Authorization", "Basic "+client.apiToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	responseBody, err := client.runRequest(req)
	if err != nil {
		return "", fmt.Errorf("%w, %w", ErrGetTokenFailed, err)
	}

	var reponseData struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	err = json.Unmarshal(responseBody, &reponseData)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGetTokenFailed, err)
	}

	return reponseData.AccessToken, nil
}

// Criteria Search Locates possible entities from the Dun & Bradstreet Data Cloud using
// specified criteria when there is not enough known data for a Match, allowing a maximum of 1000 results.
//
// # Parameters
//
// - ctx
//
// - options: variatic parameter that allows configuration of the searche and passing in the seach request body
//
// # Returns
//
// - CompanyResults: company search results
//
// - error: error if any
//
// # Documentation
//
// - https://directplus.documentation.dnb.com/openAPI.html?apiID=searchCriteria
func (client *Client) CriteriaSearch(ctx context.Context, options ...ClientOptions) (*api_response.CompanySearch, error) {
	searchResults := &api_response.CompanySearch{}
	client.RequestBody.CompanySearch = &CompanySearchRequest{}

	client.loadOptions(options...)

	reqBytes, err := json.Marshal(client.RequestBody.CompanySearch)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrSearchCriteriaFailed, err)
	}

	reqURL := client.BaseURL + CriteriaSearchURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrSearchCriteriaFailed, err)
	}

	req.Header.Add("Authorization", "Bearer "+client.apiToken)
	req.Header.Add("Content-Type", "application/json")

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrSearchCriteriaFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrSearchCriteriaFailed, err)
	}

	return searchResults, nil
}

// Tyoehead Search enables users to quickly find company records without
// having to type the entire company information in the search request.
//
// # Parameters
//
// - ctx
//
// - options: allows configuring the search
//
// - searchTerm: 2 to 30 characters used to find entities by its primary name or
// one of its tradestyle names.
//
// - countrCode: The 2-letter country/market code defined by the International Organization
// for Standardization (ISO) ISO 3166-1 scheme identifying the country of the entity.
//
// # Returns
//
// - TypeheadSearch: Typehead search company results
//
// - error: error if any
//
// # Documentation
//
// - https://directplus.documentation.dnb.com/openAPI.html?apiID=searchTypeahead
func (client *Client) TypeheadSearch(ctx context.Context, searchTerm string, countryCode string, options ...ClientOptions) (*api_response.TypeheadSearch, error) {
	searchResults := &api_response.TypeheadSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + TypeheadSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrTypeheadSearchFailed, err)
	}

	params := reqURL.Query()
	params.Add("searchTerm", searchTerm)
	params.Add("countryISOAlpha2Code", countryCode)
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrTypeheadSearchFailed, err)
	}

	req.Header.Add("Authorization", "Bearer "+client.apiToken)
	req.Header.Add("Content-Type", "application/json")

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrTypeheadSearchFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrTypeheadSearchFailed, err)
	}

	return searchResults, nil
}

// Company Lisnt Search identifies the entities from the data source that match all
// the specified criteria, allowing a maximum of 10000 results.
//
// # Parameters
//
// - ctx
//
// - options: allows configuring the search
//
// # Returns
//
// - CompanySearch: company list search results
//
// - error: error if any
//
// # Documentation
//
// - https://directplus.documentation.dnb.com/openAPI.html?apiID=searchCompanyList
func (client *Client) CompanyListSearch(ctx context.Context, options ...ClientOptions) (*api_response.CompanySearch, error) {
	searchResults := &api_response.CompanySearch{}
	client.RequestBody.CompanySearch = &CompanySearchRequest{}

	client.loadOptions(options...)

	reqBytes, err := json.Marshal(client.RequestBody.CompanySearch)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrCompanyListFailed, err)
	}

	reqURL := client.BaseURL + CompanyListURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrCompanyListFailed, err)
	}

	req.Header.Add("Authorization", "Bearer "+client.apiToken)
	req.Header.Add("Content-Type", "application/json")

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrCompanyListFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrCompanyListFailed, err)
	}

	return searchResults, nil
}

// Search Contact allows D&B Direct+ customers to search for individuals using several parameters.
// This function will perfom both standard and premium search depending on the data passed in the
// client options contact search request body. Refer to documentation for details.
//
// # Parameters
//
// - ctx
//
// - options: allows configuring the search
//
// # Returns
//
// - ContactSearch: contact search results
//
// - error: error if any
//
// # Documentation
//
// - Contact Search Standard: https://directplus.documentation.dnb.com/openAPI.html?apiID=searchContactsStandard
//
// - Contact Search Premium: https://directplus.documentation.dnb.com/openAPI.html?apiID=searchContactsPremium
func (client *Client) SearchContact(ctx context.Context, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}
	client.RequestBody.ContactSearch = &ContactSearchRequest{}

	client.loadOptions(options...)

	reqBytes, err := json.Marshal(client.RequestBody.ContactSearch)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrContactSearchFailed, err)
	}

	reqUrl := client.BaseURL + ContactSearchURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewBuffer(reqBytes))
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrContactSearchFailed, err)
	}

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrContactSearchFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrContactSearchFailed, err)
	}

	return searchResults, nil
}

// GetContactsByID will return a single contact from the D&B Direct+ API based on the contact ID
//
// # Parameters
//
// - ctx
//
// - contactID:1 to 16 characters used to find entities by a unique ID assigned to the contact.
//
// - options: allows configuring the search
//
// # Returns
//
// - ContactSearch: contact search results
//
// - error: error if any
//
// # Documentation
//
// - https://directplus.documentation.dnb.com/openAPI.html?apiID=searchContactsGet
func (client *Client) GetContactByID(ctx context.Context, contactID string, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + ContactSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrGetContactsFailed, err)
	}

	params := reqURL.Query()
	params.Add("contactID", contactID)
	reqURL.RawQuery = params.Encode()

	return client.getContact(ctx, reqURL)
}

// GetContactsByEmail will return a single contact from the D&B Direct+ API based on the contact email
//
// # Parameters
//
// - ctx
//
// - email: 1 to 128 characters used to find entities by the contact's email address.
//
// # Returns
//
// - ContactSearch: contact search results
//
// - error: error if any
//
// # Documentation
//
// - https://directplus.documentation.dnb.com/openAPI.html?apiID=searchContactsGet
func (client *Client) GetContactByEmail(ctx context.Context, email string, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + ContactSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrGetContactsFailed, err)
	}

	params := reqURL.Query()
	params.Add("contactEmail", email)
	reqURL.RawQuery = params.Encode()

	return client.getContact(ctx, reqURL)
}

// GetContactsByDUNS will return a single contact from the D&B Direct+ API based on the contact DUNS
//
// # Parameters
//
// - ctx
//
// - duns: D-U-N-S number to be used for search (mandatory when Searching by D-U-N-S, otherwise optional)
//
// # Returns
//
// - ContactSearch: contact search results
//
// - error: error if any
//
// # Documentation
//
// - https://directplus.documentation.dnb.com/openAPI.html?apiID=searchContactsGetByDuns
func (client *Client) GetContactByDUNS(ctx context.Context, duns string, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + ContactSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrGetContactsFailed, err)
	}

	params := reqURL.Query()
	params.Add("duns", duns)
	reqURL.RawQuery = params.Encode()

	return client.getContact(ctx, reqURL)
}

func (client *Client) getContact(ctx context.Context, reqURl *url.URL) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURl.String(), nil)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrGetContactsFailed, err)
	}

	req.Header.Add("Authorization", "Bearer "+client.apiToken)
	req.Header.Add("Content-Type", "application/json")

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrGetContactsFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w, %w", ErrGetContactsFailed, err)
	}

	return searchResults, nil
}

func (client *Client) runRequest(req *http.Request) ([]byte, error) {

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		errorResponse := &api_response.ErrorResponse{}

		err = json.Unmarshal(body, errorResponse)
		if err != nil {
			return nil, fmt.Errorf("%w, %d", ErrRequestFailed, res.StatusCode)
		}

		if client.BaseURL == BaseURLV3 {
			return nil, fmt.Errorf("%w, %s", ErrRequestFailed, errorResponse.ErrorDescription)
		}

		if client.BaseURL == BaseURLV1 {
			return nil, fmt.Errorf("%w, %s", ErrRequestFailed, errorResponse.ErrorMessage)
		}
	}

	return body, nil
}

func (client *Client) loadOptions(options ...ClientOptions) {
	if len(options) > 0 {
		client.options = options
	}

	for _, option := range client.options {
		option(client)
	}
}
