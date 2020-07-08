package coverages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc1(t *testing.T) {
	result := Fun1()
	assert.Equal(t, result, 1)
}

func TestFunc2(t *testing.T) {
	result := Fun2()
	assert.Equal(t, result, 2)
}
