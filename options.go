package dnbclient

type ClientOptions func(*Client)

func WithBaseURL(baseURL string) ClientOptions {
	return func(client *Client) {
		client.BaseURL = baseURL
	}
}

func WithCredentials(username string, password string) ClientOptions {
	return func(client *Client) {
		client.username = username
		client.password = password
	}
}

func WithTokens(apiKey string, apiSecret string) ClientOptions {
	return func(client *Client) {
		client.ApiKey = apiKey
		client.ApiSecret = apiSecret
	}
}

func WithAPIToken(apiToken string) ClientOptions {
	return func(client *Client) {
		client.apiToken = apiToken
	}
}

func WithCompanySerchRequest(companySearch *CompanySearchRequest) ClientOptions {
	return func(client *Client) {
		client.RequestBody.CompanySearch = companySearch
	}
}

func WithContactSearchRequest(contactSearch *ContactSearchRequest) ClientOptions {
	return func(client *Client) {
		client.RequestBody.ContactSearch = contactSearch
	}
}

func WithDUNS(duns string) ClientOptions {
	return func(client *Client) {
		client.RequestBody.CompanySearch.DUNS = duns
	}
}
