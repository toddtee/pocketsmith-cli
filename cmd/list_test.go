package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAccount(t *testing.T) {
	t.Run("returns some json", func(t *testing.T) {
		want := "something"
		got := "nothing"
		assert.Equal(t, want, got, "should be equal")
	})
}
