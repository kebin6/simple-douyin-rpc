package util

import "fmt"

const defaultCode = 3

const canceledCode = 1
const invalidArgumentCode = 3
const notFoundCode = 5
const alreadyExistsCode = 6
const internalCode = 13
const abortedCode = 10
const unavailableCode = 14
const httpErrorCode = 10000

// CodeError is error with custom error code
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// DYError 抖音返回的通用错误.
type DYError struct {
	ErrCode int64  `json:"error_code"`
	ErrMsg  string `json:"description"`
}

// DYErrorExtra 抖音返回的错误额外信息.
type DYErrorExtra struct {
	ErrCode    int64  `json:"error_code"`
	ErrMsg     string `json:"description"`
	SubErrCode int64  `json:"sub_error_code"`
	SubErrMsg  string `json:"sub_description"`
	LogID      string `json:"logid"`
	Now        int64  `json:"now"`
}

// NewCodeError returns a code error
func NewCodeError(code int, msg string, a ...any) error {
	return &CodeError{Code: code, Msg: fmt.Sprintf(msg, a...)}
}

// NewCodeCanceledError returns Code Error with custom cancel error code
func NewCodeCanceledError(msg string, a ...any) error {
	return NewCodeError(canceledCode, msg, a...)
}

// NewCodeInvalidArgumentError returns Code Error with custom invalid argument error code
func NewCodeInvalidArgumentError(msg string, a ...any) error {
	return NewCodeError(invalidArgumentCode, msg, a...)
}

// NewCodeNotFoundError returns Code Error with custom not found error code
func NewCodeNotFoundError(msg string, a ...any) error {
	return NewCodeError(notFoundCode, msg, a...)
}

// NewCodeAlreadyExistsError returns Code Error with custom already exists error code
func NewCodeAlreadyExistsError(msg string, a ...any) error {
	return NewCodeError(alreadyExistsCode, msg, a...)
}

// NewCodeAbortedError returns Code Error with custom aborted error code
func NewCodeAbortedError(msg string, a ...any) error {
	return NewCodeError(abortedCode, msg, a...)
}

// NewCodeInternalError returns Code Error with custom internal error code
func NewCodeInternalError(msg string, a ...any) error {
	return NewCodeError(internalCode, msg, a...)
}

// NewCodeUnavailableError returns Code Error with custom unavailable error code
func NewCodeUnavailableError(msg string, a ...any) error {
	return NewCodeError(unavailableCode, msg, a...)
}

// NewCodeHttpError returns Code Error with custom http error code
func NewCodeHttpError(msg string, a ...any) error {
	return NewCodeError(httpErrorCode, msg, a...)
}

// NewCodeDouYinError NewCodeDYError returns Code Error with custom douyin error code
func NewCodeDouYinError(msg string, dyError DYError, extra *DYErrorExtra) error {
	if extra != nil {
		msg = "[" + msg + "] Error : errcode=%d , errmsg=%s , suberrcode=%d, suberrmsg=%s, logid=%s"
		return NewCodeHttpError(msg, dyError.ErrCode, dyError.ErrMsg, extra.SubErrCode, extra.SubErrMsg, extra.LogID)
	}
	msg = "[" + msg + "] Error : errcode=%d , errmsg=%s"
	return NewCodeHttpError(msg, dyError.ErrCode, dyError.ErrMsg)
}

func NewDefaultError(msg string, a ...any) error {
	return NewCodeError(defaultCode, msg, a...)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
