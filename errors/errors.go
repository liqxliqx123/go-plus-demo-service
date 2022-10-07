package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	pkgErr "github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

var (
	// ErrInvalidInput 參數不合法
	ErrInvalidInput = Error{Code: "400001", Message: "One of the request inputs is not valid.", HTTPStatus: http.StatusBadRequest, GRPCCode: codes.InvalidArgument}
	// ErrInvalidInputFormat 參數不合法
	ErrInvalidInputFormat = Error{Code: "400001", Message: "One of the request inputs is not valid. (%s)", HTTPStatus: http.StatusBadRequest, GRPCCode: codes.InvalidArgument}
	// ErrUnauthorized 伺服端知道用戶端的身份，需要授權以回應請求。
	ErrUnauthorized = Error{Code: "401001", Message: http.StatusText(http.StatusUnauthorized), HTTPStatus: http.StatusUnauthorized, GRPCCode: codes.Unauthenticated}
	// ErrForbidden 伺服端不知道用戶端的身份，用戶端並無訪問權限
	ErrForbidden = Error{Code: "403000", Message: http.StatusText(http.StatusForbidden), HTTPStatus: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	// ErrAccountIsDisabled 用戶端的身份已被停用
	ErrAccountIsDisabled = Error{Code: "403001", Message: "The specified account is disabled.", HTTPStatus: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	// ErrAuthenticationFailed 伺服端無法驗證用戶身份，請確認簽章格式
	ErrAuthenticationFailed = Error{Code: "403002", Message: "Server failed to authenticate the request. Make sure the value of the Authorization header is formed correctly including the signature.", HTTPStatus: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	// ErrUsernameOrPasswordIncorrect 帳號或密碼錯誤
	ErrUsernameOrPasswordIncorrect = Error{Code: "403006", Message: "Username or Password is incorrect.", HTTPStatus: http.StatusUnauthorized, GRPCCode: codes.Unauthenticated}
	// ErrResourceNotFound 查找不到資源
	ErrResourceNotFound = Error{Code: "404001", Message: "The specified resource does not exist.", HTTPStatus: http.StatusNotFound, GRPCCode: codes.NotFound}
	// ErrResourceAlreadyExists 資源已存在
	ErrResourceAlreadyExists = Error{Code: "409004", Message: "The specified resource already exists.", HTTPStatus: http.StatusConflict, GRPCCode: codes.AlreadyExists}
	// ErrInternalError 伺服端內部錯誤
	ErrInternalError = Error{Code: "500001", Message: "The server encountered an internal models.Error. Please retry the request.", HTTPStatus: http.StatusInternalServerError, GRPCCode: codes.Internal}
)

// Export origin func of errors
var (
	Unwrap       = errors.Unwrap
	Is           = errors.Is
	As           = errors.As
	Wrap         = pkgErr.Wrap
	New          = pkgErr.New
	Cause        = pkgErr.Cause
	WithMessage  = pkgErr.WithMessage
	WithMessagef = pkgErr.WithMessagef
	WithStack    = pkgErr.WithStack
)

// StackTrace returns stack frames
func StackTrace(e error) []string {
	stacktrace := fmt.Sprintf("%+v\n", e)
	output := strings.Split(stacktrace, "\n")
	return output[:len(output)-1]
}
