package dnbclient

type ClientOptions func(*Client)

func WithBaseURL(baseURL string) ClientOptions {
	return func(client *Client) {
		client.BaseURL = baseURL
	}
}

func WithAPIKey(apiKey string) ClientOptions {
	return func(client *Client) {
		client.apiToken = apiKey
	}
}

func WithSearchCriteriaRequest(body *RequestBody) ClientOptions {
	return func(client *Client) {
		client.RequestBody = body
	}
}

func WithCompanyListRequest(body *RequestBody) ClientOptions {
	return func(client *Client) {
		client.RequestBody = body
	}
}
