package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrivial(t *testing.T) {
	assert.Equal(t, 1+1, 2, "This computer is broken.")
}
