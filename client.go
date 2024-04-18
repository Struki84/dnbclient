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

	"github.com/struki84/dnbclient/api_response"
)

const (
	DefaultBaseURL    = "https://plus.dnb.com/v1"
	AuthURL           = "/token"
	CriteriaSearchURL = "/search/criteria"
	TypeheadSearchURL = "/search/typehead"
	CompanyListURL    = "/search/companyList"
	ContactSearchURL  = "/search/contact"
)

var (
	ErrMissingAPIKey        = errors.New("api token is required")
	ErrGetTokenFailed       = errors.New("get token failed with error")
	ErrRequestFailed        = errors.New("http request failed with error")
	ErrSearchCriteriaFailed = errors.New("search criteria failed with error")
	ErrCompanyListFailed    = errors.New("company list search failed with error")
	ErrTypeheadSearchFailed = errors.New("typehead search failed with error")
	ErrContactSearchFailed  = errors.New("contact search failed with error")
	ErrGetContactsFailed    = errors.New("get contacts failed with error")
	ErrNoSearchResults      = errors.New("no search results found")
)

type Client struct {
	username    string
	password    string
	apiToken    string
	options     []ClientOptions
	BaseURL     string
	RequestBody *RequestBody
}

func NewClient(apiToken string, options ...ClientOptions) (*Client, error) {
	if apiToken == "" {
		return nil, ErrMissingAPIKey
	}

	client := &Client{
		apiToken:    apiToken,
		BaseURL:     DefaultBaseURL,
		RequestBody: &RequestBody{},
	}

	return client, nil
}

func (client *Client) GetToken(ctx context.Context, options ...ClientOptions) (string, error) {
	var reqData struct {
		GrantType string `json:"grant_type"`
	}
	reqData.GrantType = "client_credentials"

	client.loadOptions(options...)

	credentials := client.username + ":" + client.password
	client.apiToken = base64.StdEncoding.EncodeToString([]byte(credentials))

	reqBytes, err := json.Marshal(reqData)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGetTokenFailed, err)
	}

	reqURL := client.BaseURL + AuthURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGetTokenFailed, err)
	}

	responseBody, err := client.runRequest(req)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGetTokenFailed, err)
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

func (client *Client) CriteriaSearch(ctx context.Context, options ...ClientOptions) (*api_response.CompanySearch, error) {
	searchResults := &api_response.CompanySearch{}
	client.RequestBody.CompanySearch = &CompanySearchRequest{}

	client.loadOptions(options...)

	reqBytes, err := json.Marshal(client.RequestBody.CompanySearch)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrSearchCriteriaFailed, err)
	}

	reqURL := client.BaseURL + CriteriaSearchURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrSearchCriteriaFailed, err)
	}

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrSearchCriteriaFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrSearchCriteriaFailed, err)
	}

	return searchResults, nil
}

func (client *Client) TypeheadSearch(ctx context.Context, searchTerm string, countryCode string, options ...ClientOptions) (*api_response.TypeheadSearch, error) {
	searchResults := &api_response.TypeheadSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + TypeheadSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrTypeheadSearchFailed, err)
	}

	params := reqURL.Query()
	params.Add("searchTerm", searchTerm)
	params.Add("countryISOAlpha2Code", countryCode)
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrTypeheadSearchFailed, err)
	}

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrTypeheadSearchFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrTypeheadSearchFailed, err)
	}

	return searchResults, nil
}

func (client *Client) CompanyListSearch(ctx context.Context, options ...ClientOptions) (*api_response.CompanySearch, error) {
	searchResults := &api_response.CompanySearch{}
	client.RequestBody.CompanySearch = &CompanySearchRequest{}

	client.loadOptions(options...)

	reqBytes, err := json.Marshal(client.RequestBody.CompanySearch)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrCompanyListFailed, err)
	}

	reqURL := client.BaseURL + CompanyListURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrCompanyListFailed, err)
	}

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrCompanyListFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrCompanyListFailed, err)
	}

	return searchResults, nil
}

func (client *Client) SearchContact(ctx context.Context, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}
	client.RequestBody.ContactSearch = &ContactSearchRequest{}

	client.loadOptions(options...)

	reqBytes, err := json.Marshal(client.RequestBody.ContactSearch)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrContactSearchFailed, err)
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

func (client *Client) GetContactByID(ctx context.Context, contactID string, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + ContactSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrGetContactsFailed, err)
	}

	params := reqURL.Query()
	params.Add("contactID", contactID)
	reqURL.RawQuery = params.Encode()

	return client.getContact(ctx, reqURL)
}

func (client *Client) GetContactByEmail(ctx context.Context, email string, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + ContactSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrGetContactsFailed, err)
	}

	params := reqURL.Query()
	params.Add("contactEmail", email)
	reqURL.RawQuery = params.Encode()

	return client.getContact(ctx, reqURL)
}

func (client *Client) GetcontactByDUNS(ctx context.Context, duns string, options ...ClientOptions) (*api_response.ContactSearch, error) {
	searchResults := &api_response.ContactSearch{}

	client.loadOptions(options...)

	reqURL, err := url.Parse(client.BaseURL + ContactSearchURL)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrGetContactsFailed, err)
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
		return searchResults, fmt.Errorf("%w: %w", ErrGetContactsFailed, err)
	}

	responseBody, err := client.runRequest(req)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrGetContactsFailed, err)
	}

	err = json.Unmarshal(responseBody, searchResults)
	if err != nil {
		return searchResults, fmt.Errorf("%w: %w", ErrGetContactsFailed, err)
	}

	return searchResults, nil
}

func (client *Client) runRequest(req *http.Request) ([]byte, error) {

	req.Header.Add("Authorization", "Bearer "+client.apiToken)
	req.Header.Add("Content-Type", "application/json")

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
			return nil, fmt.Errorf("%w: %d", ErrRequestFailed, res.StatusCode)
		}

		return nil, fmt.Errorf("%w: %s", ErrRequestFailed, errorResponse.ErrorMessage)
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
