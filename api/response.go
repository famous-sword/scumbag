package api

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	Success = NewResponse(http.StatusOK, "ok")
	Error   = NewResponse(http.StatusInternalServerError, "error")
)

func (r *Response) WithMessage(message string) *Response {
	r.Message = message

	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data

	return r
}

func NewResponse(code int, message string) *Response {
	return &Response{Code: code, Message: message}
}
