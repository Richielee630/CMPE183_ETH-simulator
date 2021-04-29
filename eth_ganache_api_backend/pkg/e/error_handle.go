package e

import "net/http"

// api错误的结构体
type APIException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}

const (
	SERVER_ERROR = 1000 // 系统错误

	NOT_FOUND = 1001 // 401错误

	UNKNOWN_ERROR = 1002 // 未知错误

	PARAMETER_ERROR = 1003 // 参数错误

	AUTH_ERROR = 1004 // 错误

)

// 500 错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, SERVER_ERROR, http.StatusText(http.StatusInternalServerError))
}

// 404 错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, UNKNOWN_ERROR, message)
}

// 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, PARAMETER_ERROR, message)
}
