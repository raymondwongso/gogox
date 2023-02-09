package trace

import (
	"net/http"

	"github.com/raymondwongso/gogox/log"
	"github.com/raymondwongso/gogox/trace"
)

// TraceMiddleware provides middleware to inject traceID
func TraceMiddleware(traceField string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		traceID := trace.TraceFromContext(ctx)
		if traceID == "" {
			traceID = trace.New()
		}

		ctx = trace.NewContext(ctx, traceID)

		logMd := log.MetadataFromContext(ctx)
		logMd[traceField] = traceID

		ctx = log.NewContext(ctx, logMd)

		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}
