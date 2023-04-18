package log_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raymondwongso/gogox/http/log"
	"github.com/raymondwongso/gogox/log/nop"
	"github.com/stretchr/testify/assert"
)

// TODO(raymondwongso): looks for better testing method
func Test_LoggingMiddleware(t *testing.T) {
	t.Run("success should log", func(t *testing.T) {
		assert.NotPanics(t, func() {
			h := log.LoggingMiddleware(nop.New(), http.DefaultServeMux, log.DefaultOptions())
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
		})
	})

	t.Run("success should not log", func(t *testing.T) {
		assert.NotPanics(t, func() {
			opts := log.DefaultOptions()
			opts.ShouldLog = func(req *http.Request, status int) bool {
				return false
			}

			h := log.LoggingMiddleware(nop.New(), http.DefaultServeMux, opts)
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
		})
	})

}
