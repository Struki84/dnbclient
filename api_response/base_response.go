package api_response

type Base struct {
	TransactionDetail          TransactionDetail `json:"transactionDetail,omitempty"`
	CandidatesMatchedQuantity  int               `json:"candidatesMatchedQuantity,omitempty"`
	CandidatesReturnedQuantity int               `json:"candidatesReturnedQuantity,omitempty"`
}

type TransactionDetail struct {
	TransactionID        string `json:"transactionID,omitempty"`
	TransactionTimestamp string `json:"transactionTimestamp,omitempty"`
	InLanguage           string `json:"inLanguage,omitempty"`
	ServiceVersion       string `json:"serviceVersion,omitempty"`
}

type ErrorResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
