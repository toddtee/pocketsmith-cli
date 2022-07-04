package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	testCases := map[string]struct {
		path string
	}{
		"Everything OK": {
			path: "/me",
		},
	}
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				rp := r.URL.Path
				if rp != tt.path {
					t.Errorf("Expected request path %v, got %v", tt.path, rp)
				}
				//		w.Write()
			}))
			defer svr.Close()
			cfg := BuildConfig()
			cfg.BaseURL = svr.URL

			AuthorisedUser(cfg)

			// assert.Equal(t, tt.path, r.URL.Path, "should be equal")
		})
	}
}
