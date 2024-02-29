package response

import (
	"fmt"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeValue  = "application/json"
	messageKey        = "message"
)

func NewErrorResponse(statusCode int, msg string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Header().Set(contentTypeHeader, contentTypeValue)
	_, _ = w.Write([]byte(fmt.Sprintf("{\"%s\": \"%s\"} \n", messageKey, msg)))
}
