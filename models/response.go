package models

import (
	"encoding/json"
	"fmt"
)

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  ErrorList   `json:"errors"`
}

type Error struct {
	HTTPStatus int         `json:"-"`
	Message    string      `json:"message,omitempty"`
	Code       string      `json:"code,omitempty"`
	ErrorData  interface{} `json:"error_data,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("message:%s-code:%s", e.Message, e.Code)
}

// ErrorWithData ErrorWithData
type ErrorWithData struct {
	Error
	Data interface{} `json:"data"`
}

// ErrorWithDataList error list struct
type ErrorWithDataList []ErrorWithData

// AddError AddError
func (e *ErrorWithDataList) AddError(err error, code string, data interface{}) {
	*e = append(*e, ErrorWithData{
		Error: Error{
			Message: err.Error(),
			Code:    code,
		},
		Data: data,
	})
}

// AppendError AppendError
func (e *ErrorWithDataList) AppendError(err Error, data interface{}) {
	*e = append(*e, ErrorWithData{
		Error: err,
		Data:  data,
	})
}

// AddIfExistErrorCode500 ErrorList add error if exist
func (e *ErrorWithDataList) AddIfExistErrorCode500(err error, msg string) {
	if err != nil {
		*e = append(*e, ErrorWithData{
			Error: Error{
				Message: fmt.Sprintf("%s:%s", msg, err.Error()),
				Code:    "500",
			}})
	}

}

// ErrorList error list struct
type ErrorList []Error

// AddError ErrorList add error
func (e *ErrorList) AddError(err error, code string) {
	*e = append(*e, Error{
		Message: err.Error(),
		Code:    code,
	})
}

// AddErrorWithErrorData Add Error With ErrorData
func (e *ErrorList) AddErrorWithErrorData(err error, code string, errorData interface{}) {
	*e = append(*e, Error{
		Message:   err.Error(),
		Code:      code,
		ErrorData: errorData,
	})
}

// AddIfExistErrorCode500 ErrorList add error if exist
func (e *ErrorList) AddIfExistErrorCode500(err error, msg string) {
	if err != nil {
		*e = append(*e, Error{
			Message: fmt.Sprintf("%s:%s", msg, err.Error()),
			Code:    "500",
		})
	}

}

// ToJSON ToJSON
func (e *ErrorList) ToJSON() json.RawMessage {
	marshal, _ := json.Marshal(e)
	return marshal
}

// ToError error to []Error
func ToError(err error, code string) []Error {
	return []Error{
		{
			Message: err.Error(),
			Code:    code,
		},
	}
}
