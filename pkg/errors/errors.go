package errors

import (
	"encoding/json"
	"io"
)

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponseData []ErrorData

func (er *ErrorResponseData) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(er)
}
