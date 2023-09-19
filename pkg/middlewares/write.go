package middlewares

import "net/http"

func WriteResponse(r *http.ResponseWriter, statusCode int, data interface{}) {
	(*r).Header().Set("Content-Type", "application/json")
	(*r).WriteHeader(statusCode)
	_, _ = (*r).Write(data.([]byte))
}
