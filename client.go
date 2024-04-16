package dnbclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	DefaultBaseURL    = "https://plus.dnb.com/v1"
	CriteriaSearchURL = "/search/criteria"
	TypeheadSearchURL = "/search/typehead"
	CompanyListURL    = "/search/companyList"
)

var (
	ErrMissingAPIKey        = errors.New("api token is required")
	ErrRequestFailed        = errors.New("request failed with error")
	ErrSearchCriteriaFailed = errors.New("search criteria failed with error")
	ErrCompanyListFailed    = errors.New("company list search failed with error")
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

func (client *Client) SearchCriteria(ctx context.Context, options ...ClientOptions) (*CriteriaSearchResponse, error) {
	searchResults := &CriteriaSearchResponse{}
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

func (client *Client) CompanyList(ctx context.Context, options ...ClientOptions) (*CompanyListResponse, error) {
	searchResults := &CompanyListResponse{}
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
		return nil, fmt.Errorf("%w: %d", ErrRequestFailed, res.StatusCode)
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
