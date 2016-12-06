package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueryKey(t *testing.T) {
	name := GetNameQueryKey()
	assert.Equal(t, "name", name, "incorrect key")
}
