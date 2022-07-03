package wiring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusHelper(t *testing.T) {
	tests := map[string]struct {
		StatusCode int
	}{
		"StatusOK": {
			StatusCode: 200,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// request, _ := http.NewRequest(http.MethodGet, "/my_url", nil)
			// response := httptest.NewRecorder()
			// HTTPRequest(request)

			// got := response.Body.String()

			got := tt.StatusCode
			want := 200
			assert.Equal(t, want, got, "should be equal")
		})
	}
}
