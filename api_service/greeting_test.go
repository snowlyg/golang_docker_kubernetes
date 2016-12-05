package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	name := "Boris"

	greeting := ConstructGreeting(name)
	assert.Equal(t, "Hello Boris from Distelli!", greeting, "incorrect greeting")
}
