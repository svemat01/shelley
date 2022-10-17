package pkg

import "encoding/json"

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

func ValidatioError(errors []*ErrorResponse) *APIErrorData {
	return &APIErrorData{Status: 400, Code: "validation-error", Data: errors}
}

func Unexpected(m string) *APIError {
	return &APIError{Status: 500, Code: "internal-server", Message: m}
}

// --------------------------------------------

type APIErrorData struct {
	Status int         `json:"status"`
	Code   string      `json:"code"`
	Data   interface{} `json:"data"`
}

func (e *APIErrorData) Error() string {
	data, err := json.Marshal(e.Data)

	if err != nil {
		return err.Error()
	}

	return string(data)
}
