package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestAddOK(t *testing.T) {
	assert.Equal(t, 3, Add(1, 2))
}

func TestAddFail(t *testing.T) {
	assert.Equal(t, 2, Add(1, 1))
}
