package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (w *resLoggingWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()
		log.Printf("[%d]%s %s", traceID, req.Method, req.URL.Path)

		ctx := SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.statusCode)
	})
}
