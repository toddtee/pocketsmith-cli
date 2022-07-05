package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := map[string]struct {
		path string
		user string
	}{
		"Get Autorised User": {
			path: "/me",
			user: "Monty Burns",
		},
	}
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				rp := r.URL.Path
				if rp != tt.path {
					t.Errorf("Expected request path %v, got %v", tt.path, rp)
				}
				w.Write([]byte(tt.user))
			}))
			defer svr.Close()
			cfg := BuildConfig()
			cfg.BaseURL = svr.URL

			// resp := GetAuthorisedUser()
			// d, _ := ioutil.ReadAll(resp.Body)
			got := "hello"

			assert.Equal(t, tt.user, got, "should be equal")
		})
	}
}
