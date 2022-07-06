package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorisedUser(t *testing.T) {
	testCases := map[string]struct {
		Path       string
		User       User
		StatusCode int
	}{
		"HappyPath": {
			Path: "/me",
			User: User{
				Name:  "Monty Burns",
				Email: "mburns@springfieldnukepower.com",
			},
			StatusCode: http.StatusOK,
		},
		"HappyPathNoName": {
			Path: "/me",
			User: User{
				Email: "mburns@springfieldnukepower.com",
			},
			StatusCode: http.StatusOK,
		},
		"HappyPathNoEmail": {
			Path: "/me",
			User: User{
				Name: "Monty Burns",
			},
			StatusCode: http.StatusOK,
		},
	}
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			// Create a test server
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				rp := r.URL.Path
				// Check the path is as expected
				if rp != tt.Path {
					t.Errorf("Expected request path %v, got %v", tt.Path, rp)
				}
				pl := fmt.Sprintf("{\"name\":\"%s\", \"email\":\"%s\"}", tt.User.Name, tt.User.Email)
				// Write the response body
				w.Write([]byte(pl))
				// Write the response header
				w.WriteHeader(tt.StatusCode)
			}))

			// Make a new client and override the BaseURL to be the test server's URL
			c := NewClient()
			c.BaseURL = svr.URL
			// Call the function under test
			u, _ := c.GetAuthorisedUser()
			defer svr.Close()

			assert.Equal(t, tt.User.Name, u.Name, "Name of User should match.")
			assert.Equal(t, tt.User.Email, u.Email, "Email of User should match.")
		})
	}
}
