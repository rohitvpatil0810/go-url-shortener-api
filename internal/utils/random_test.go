package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	tests := []int{5, 10, 15, 20}
	for _, n := range tests {
		t.Run("Running for "+strconv.Itoa(n), func(t *testing.T) {
			t.Parallel()
			s := RandomString(n)
			assert.Equal(t, n, len(s), "Length should be same")
		})
	}
}

func TestRandomInt(t *testing.T) {
	min, max := 10, 20
	n := RandomInt(min, max)
	assert.GreaterOrEqual(t, n, min, "Should be greater or equal to min")
	assert.LessOrEqual(t, n, max, "Should be less or equal to max")
}

func TestRandomUsername(t *testing.T) {
	s := RandomUsername()
	assert.Equal(t, 7, len(s), "Length should be 7")
}

func TestRandomPassword(t *testing.T) {
	s := RandomPassword()
	assert.Equal(t, 10, len(s), "Length should be 10")
}
