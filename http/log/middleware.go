package log

import (
	"net/http"
	"time"

	"github.com/raymondwongso/gogox/log"
)

type loggingResponseWriter struct {
	responseWriter http.ResponseWriter
	status         int
	size           int
}

func (r *loggingResponseWriter) Header() http.Header {
	return r.responseWriter.Header()
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.responseWriter.Write(b)
	r.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.responseWriter.WriteHeader(statusCode)
	r.status = statusCode
}

// LoggingMiddleware provides log middleware that log request and response.
func LoggingMiddleware(logger log.Logger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		md := newLogMetadata(r, startTime)

		lrw := &loggingResponseWriter{
			responseWriter: w,
		}

		handler.ServeHTTP(lrw, r)

		duration_ms := float32(time.Since(startTime).Nanoseconds()/1000) / 1000
		size := lrw.size
		status := lrw.status

		// TODO(raymondwongso) research more about status in integration with grpc gateway, somehow status is 0
		// for now, override to 200 assuming it is a successful request
		if status == 0 {
			status = http.StatusOK
		}

		md["http.duration_ms"] = duration_ms
		md["http.response_size"] = size
		md["http.response_status"] = status

		logLevel := log.DebugLevel
		if status >= 500 {
			logLevel = log.ErrorLevel
		} else if status >= 400 {
			logLevel = log.WarnLevel
		} else {
			logLevel = log.InfoLevel
		}

		logger.Logw(logLevel, "[%s]%s finished http request with status: %d", md, r.Method, r.URL.Path, status)
	})
}

func newLogMetadata(r *http.Request, startTime time.Time) log.Metadata {
	ctx := r.Context()
	logMd := log.MetadataFromContext(ctx)
	logMd["http.start_time"] = startTime.Format(time.RFC3339)

	logMd["http.header"] = r.Header
	logMd["http.method"] = r.Method
	logMd["http.url_path"] = r.URL.Path
	logMd["http.content_type"] = r.Header.Get("Content-Type")
	logMd["http.accept_encoding"] = r.Header.Get("Accept-Encoding")
	logMd["http.content_length"] = r.ContentLength
	logMd["http.user_agent"] = r.UserAgent()
	logMd["http.host"] = r.Host
	logMd["http.remote_addr"] = r.RemoteAddr

	if d, ok := ctx.Deadline(); ok {
		logMd["http.request.deadline"] = d.Format(time.RFC3339)
	}

	return logMd
}
