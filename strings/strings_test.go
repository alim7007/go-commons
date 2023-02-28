package strings_test

import (
	"testing"

	"github.com/alim7007/go-commons/strings"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Run("normal string, success", func(t *testing.T) {
		input := "Olim"
		expected := "milO"

		got := strings.Reverse(input)
		assert.Equal(t, expected, got)
	})
}
