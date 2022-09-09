package dto

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/errors"
)

type BaseResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Errors  errors.ErrorResponseData `json:"error"`
	Data    interface{}              `json:"data"`
}

func (br *BaseResponse) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(br)
}

func (br *BaseResponse) SendResponse(rw *http.ResponseWriter) error {
	(*rw).WriteHeader(br.Code)
	return json.NewEncoder(*rw).Encode(br)
}
