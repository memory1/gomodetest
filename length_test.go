package gomodtest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMile(t *testing.T) {
	assert.True(t, Mile(3).Equal(Mile(3)))
	assert.True(t, Mile(3).NotEqual(Mile(2)))
}

func TestMileAndYard(t *testing.T) {
	assert.True(t, Mile(1).Equal(Yard(1760)))
	assert.True(t, Mile(1).NotEqual(Yard(1761)))

	assert.True(t, Yard(3).Equal(Yard(3)))
	assert.True(t, Yard(3).NotEqual(Yard(4)))
}
