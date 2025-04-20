package errors

import (
	"errors"
	"net/http"
)

// 错误码
const (
	CodeInvalidRequest = "INVALID_REQUEST"
	CodeUnauthorized   = "UNAUTHORIZED"
	CodeForbidden      = "FORBIDDEN"
	CodeNotFound       = "NOT_FOUND"
	CodeInternalError  = "INTERNAL_ERROR"
)

// Error 自定义错误类型
type Error struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

// Error 实现 error 接口
func (e *Error) Error() string {
	return e.Message
}

// New 创建新的错误
func New(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// WithDetails 添加错误详情
func (e *Error) WithDetails(details map[string]string) *Error {
	e.Details = details
	return e
}

// HTTPStatus 获取 HTTP 状态码
func (e *Error) HTTPStatus() int {
	switch e.Code {
	case CodeInvalidRequest:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeInternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// 预定义错误
var (
	ErrInvalidRequest = New(CodeInvalidRequest, "无效的请求")
	ErrUnauthorized   = New(CodeUnauthorized, "未授权")
	ErrForbidden      = New(CodeForbidden, "禁止访问")
	ErrNotFound       = New(CodeNotFound, "资源不存在")
	ErrInternalError  = New(CodeInternalError, "服务器内部错误")
)

// Is 判断错误是否匹配
func Is(err error, target *Error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Code == target.Code
	}
	return false
}
