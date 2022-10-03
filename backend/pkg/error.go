package pkg

type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NotFound(m string) *APIError {
	return &APIError{Status: 404, Code: "not-found", Message: m}
}

func BadRequest(m string) *APIError {
	return &APIError{Status: 400, Code: "bad-request", Message: m}
}

func Unexpected(m string) *APIError {
	return &APIError{Status: 500, Code: "internal-server", Message: m}
}
