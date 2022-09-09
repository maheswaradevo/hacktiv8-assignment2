package utils

import (
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/dto"
	er "github.com/maheswaradevo/hacktiv8-assignment2/pkg/errors"
)

func NewErrorResponseValue(code string, value string) er.ErrorData {
	return er.ErrorData{
		Code:    code,
		Message: value,
	}
}

func NewErrorResponseData(errorResponsesValue ...er.ErrorData) er.ErrorResponseData {
	errors := er.ErrorResponseData{}

	for _, v := range errorResponsesValue {
		errors = append(errors, v)
	}
	return errors
}

func NewErrorResponse(code int, message string, errors ...er.ErrorData) *dto.BaseResponse {
	return &dto.BaseResponse{
		Code:    code,
		Message: message,
		Errors: NewErrorResponseData(
			errors...,
		),
		Data: nil,
	}
}

func NewBaseResponse(code int, message string, errors er.ErrorResponseData, data interface{}) *dto.BaseResponse {
	return &dto.BaseResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}
