package trace

import (
	"net/http"

	"github.com/raymondwongso/gogox/log"
	"github.com/raymondwongso/gogox/trace"
)

// TraceMiddleware provides middleware to inject traceID
// traceField is a log field for storing the traceID
// traceHeaderKey is a HTTP header key that contains request traceID (if provided by client)
func TraceMiddleware(traceField, traceHeaderKey string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// get traceID from request header, if not found check request context.
		// otherwise, generate new trace
		traceID := r.Header.Get(traceHeaderKey)
		if traceID == "" {
			traceID = trace.TraceFromContext(ctx)
			if traceID == "" {
				traceID = trace.New()
			}
		}

		// we set traceID to context AND also set to request header. Cause apparently grpc gateway doesn't propagate the ctx?
		ctx = trace.NewContext(ctx, traceID)
		r.Header.Set(traceHeaderKey, traceID)

		logMd := log.MetadataFromContext(ctx)
		logMd[traceField] = traceID

		ctx = log.NewContext(ctx, logMd)

		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}
