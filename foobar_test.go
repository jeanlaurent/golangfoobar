package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneIs1(t *testing.T) {
	assert.Equal(t, Foobar(1), "1")
}

func TestTwoIs2(t *testing.T) {
	assert.Equal(t, Foobar(2), "2")
}

func TestThreeIsFoo(t *testing.T) {
	assert.Equal(t, Foobar(3), "foo")
}

func TestFourIsBar(t *testing.T) {
	assert.Equal(t, Foobar(5), "bar")
}

func TestNineIsFoo(t *testing.T) {
	assert.Equal(t, Foobar(9), "foo")
}

func TestTenIsBar(t *testing.T) {
	assert.Equal(t, Foobar(10), "bar")
}

func TestFifteenIsFooBar(t *testing.T) {
	assert.Equal(t, Foobar(15), "foobar")
}
