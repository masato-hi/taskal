package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimTailingSpace(t *testing.T) {
	t.Run("When the starting character is a space.", func(t *testing.T) {
		assert := assert2.New(t)
		str := "  Hello World!"
		actual := TrimTailingSpace(str)
		expected := "  Hello World!"
		assert.Equal(expected, actual)
	})

	t.Run("When the tailing character is a space.", func(t *testing.T) {
		assert := assert2.New(t)
		str := "Hello World!  "
		actual := TrimTailingSpace(str)
		expected := "Hello World!"
		assert.Equal(expected, actual)
	})
}

func TestQuoteString(t *testing.T) {
	t.Run("When quotes are not included.", func(t *testing.T) {
		assert := assert2.New(t)
		str := "Hello World!"
		actual := QuoteString(str)
		expected := "\"Hello World!\""
		assert.Equal(expected, actual)
	})

	t.Run("When single-quotes are included.", func(t *testing.T) {
		assert := assert2.New(t)
		str := "'Hello World!'"
		actual := QuoteString(str)
		expected := "\"'Hello World!'\""
		assert.Equal(expected, actual)
	})

	t.Run("When double-quotes are included.", func(t *testing.T) {
		assert := assert2.New(t)
		str := "\"Hello World!\""
		actual := QuoteString(str)
		expected := "\"\\\"Hello World!\\\"\""
		assert.Equal(expected, actual)
	})
}
