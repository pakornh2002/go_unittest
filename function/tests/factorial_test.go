package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"

	factorial "pakornh2002/go_unittest/function"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 2", 2, 2},
		{"Factorial of 3", 3, 6},
		{"Factorial of 4", 4, 24},
		{"Factorial of 5", 5, 120},
		{"Factorial of 6", 6, 720},
		{"Factorial of -1 (negative number)", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := factorial.Factorial(tt.input)
			assert.Equal(t, tt.expected, result, "they should be equal")
		})
	}
}
