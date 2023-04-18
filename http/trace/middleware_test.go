package trace_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raymondwongso/gogox/http/trace"
	"github.com/stretchr/testify/assert"
)

// TODO(raymondwongso): looks for better testing method
func Test_TraceMiddlware(t *testing.T) {
	assert.NotPanics(t, func() {
		h := trace.TraceMiddleware("traceField", "X-Trace-Header-Key", http.DefaultServeMux)
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
	})
}
