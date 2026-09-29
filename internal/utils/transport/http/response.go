package transport_http_utils

import (
	"encoding/json"
	"net/http"
)

type HTTPResponceHandler struct {
	w http.ResponseWriter
}

type ApiError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Responce struct {
	Success any       `json:"success,omitempty"`
	Error   *ApiError `json:"error,omitempty"`
}

func NewHTTPResponceHandler(w http.ResponseWriter) HTTPResponceHandler {
	return HTTPResponceHandler{
		w: w,
	}
}

func (rw *HTTPResponceHandler) Success(status int, data any) {
	rw.Json(status, Responce{
		Success: data,
	})
}

func (rw *HTTPResponceHandler) Error(status int, code string, message string) {
	rw.Json(status, Responce{
		Error: &ApiError{
			Code:    code,
			Message: message,
		},
	})
}

func (rw *HTTPResponceHandler) Json(status int, responce any) {
	rw.w.Header().Set("Content-Type", "application/json")
	rw.w.WriteHeader(status)
	json.NewEncoder(rw.w).Encode(responce)
}

func (rw *HTTPResponceHandler) SetCookie(coockies []*http.Cookie) {
	for _, coockie := range coockies {
		http.SetCookie(rw.w, coockie)
	}
}
