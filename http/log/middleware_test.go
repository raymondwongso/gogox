package log_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	httplog "github.com/raymondwongso/gogox/http/log"
	"github.com/raymondwongso/gogox/log"
	logmock "github.com/raymondwongso/gogox/log/mock"
	"github.com/raymondwongso/gogox/log/nop"
	"github.com/stretchr/testify/assert"
)

type mdHeaderCustomMatcher struct {
	keyThatShouldExists    []string
	keyThatShouldNotExists []string
	header                 http.Header
	debval                 interface{}
}

// Matches returns whether x is a match.
func (m *mdHeaderCustomMatcher) Matches(x interface{}) bool {
	md, ok := x.(log.Metadata)
	if !ok {
		return false
	}

	header, exists := md["http.header"]
	if !exists {
		return false
	}

	httpheader, ok := header.(http.Header)
	if !ok {
		return false
	}

	m.header = httpheader
	for _, key := range m.keyThatShouldExists {
		val := httpheader.Get(key)
		if val == "" {
			return false
		}
	}

	for _, key := range m.keyThatShouldNotExists {
		val := httpheader.Get(key)
		if val != "" {
			return false
		}
	}

	return true
}

// String describes what the matcher matches.
func (m *mdHeaderCustomMatcher) String() string {
	return fmt.Sprintf("expected md has key %v and doesnt have key %v. got : %v", m.keyThatShouldExists, m.keyThatShouldNotExists, m.header.Get("ValidKey"))
}

// TODO(raymondwongso): looks for better testing method
func Test_LoggingMiddleware(t *testing.T) {
	t.Run("success should log", func(t *testing.T) {
		assert.NotPanics(t, func() {
			h := httplog.LoggingMiddleware(nop.New(), http.DefaultServeMux, httplog.DefaultOptions())
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
		})
	})

	t.Run("success should not log", func(t *testing.T) {
		assert.NotPanics(t, func() {
			opts := httplog.DefaultOptions()
			opts.ShouldLog = func(req *http.Request, status int) bool {
				return false
			}

			h := httplog.LoggingMiddleware(nop.New(), http.DefaultServeMux, opts)
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
		})
	})

	t.Run("success excluding header from log", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := logmock.NewMockLogger(ctrl)
		opts := httplog.DefaultOptions()
		opts = httplog.AddExcludedHeaderKey(opts, []string{"Authorization"})

		h := httplog.LoggingMiddleware(logger, http.DefaultServeMux, opts)
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Add("Authorization", "Bearer sometoken")
		req.Header.Add("ValidKey", "Not Deleted")
		w := httptest.NewRecorder()

		mdMatcher := mdHeaderCustomMatcher{
			keyThatShouldExists:    []string{"ValidKey"},
			keyThatShouldNotExists: []string{"Authorization"},
		}
		logger.EXPECT().Logw(gomock.Any(), "[%s]%s finished http request with status: %d", &mdMatcher, http.MethodGet, "/test", gomock.Any())

		h.ServeHTTP(w, req)
	})

}
