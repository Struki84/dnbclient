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

func WithAPIKey(apiKey string) ClientOptions {
	return func(client *Client) {
		client.apiToken = apiKey
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
