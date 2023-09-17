package middlewares

import (
	"encoding/json"
	"logur.dev/logur"
	"net/http"
	"time"
)

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	TimeStamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	ResponseData interface{} `json:"message"`
	statusCode   int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{
		w,
		nil,
		http.StatusOK,
	}
}
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
func (lrw *LoggingResponseWriter) Write(data []byte) (int, error) {
	err := json.Unmarshal(data, &lrw.ResponseData)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func Middleware(logger logur.LoggerFacade) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("--> Request: ", map[string]interface{}{
				"method": r.Method,
				"path":   r.URL.Path,
			})

			lrw := NewLoggingResponseWriter(w)

			next.ServeHTTP(lrw, r)
			statusCode := lrw.statusCode
			logger.Info("<-- Status: ", map[string]interface{}{
				"status":      statusCode,
				"status_text": http.StatusText(statusCode),
			})
			resData := Response{
				Status:    statusCode,
				Message:   http.StatusText(statusCode),
				TimeStamp: time.Now().Unix(),
				Data:      lrw.ResponseData,
			}
			responseJson, err := json.Marshal(resData)
			if err != nil {
				logger.Error("Error while marshal response: ", map[string]interface{}{
					"error": err.Error(),
				})
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(responseJson)
		})
	}
}
