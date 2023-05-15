package schema

type ErrorResult struct {
	Status bool      `json:"status"`
	Error  ErrorItem `json:"error"`
}

type ErrorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
