package factorial

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	factorial "pakornh2002/go_unittest/function"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{3, 6},
	}

	for _, tt := range tests {
		t.Run("It should return correct int answer for input "+fmt.Sprint(tt.input), func(t *testing.T) {
			result := factorial.Factorial(tt.input)
			assert.Equal(t, tt.expected, result, "they should be equal")
		})
	}
}
