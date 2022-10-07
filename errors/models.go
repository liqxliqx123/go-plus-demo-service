package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

// Error is the struct of Error segment
type Error struct {
	HTTPStatus int        `json:"-"`
	Message    string     `json:"message"`
	Code       string     `json:"code"`
	GRPCCode   codes.Code `json:"grpccode"`
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// WithMessagef annotates err with the format specifier.
func (e Error) WithMessagef(args ...interface{}) Error {
	return Error{
		HTTPStatus: e.HTTPStatus,
		Message:    fmt.Sprintf(e.Message, args...),
		Code:       e.Code,
		GRPCCode:   e.GRPCCode,
	}
}
