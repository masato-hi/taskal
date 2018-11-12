package main

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPrintf(t *testing.T) {
	assert := assert.New(t)
	buffer := &bytes.Buffer{}
	Stdout = buffer
	color.NoColor = false

	t.Run("When not passing format string", func(t *testing.T) {
		buffer.Reset()

		Printf("Test Printf")
		expected := "Test Printf\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})

	t.Run("When passing format string", func(t *testing.T) {
		buffer.Reset()

		Printf("Test %s", "Printf2")
		expected := "Test Printf2\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})
}

func TestDebug(t *testing.T) {
	assert := assert.New(t)
	buffer := &bytes.Buffer{}
	Stdout = buffer
	color.NoColor = false
	Now = func() time.Time {
		return time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	}

	t.Run("When not passing format string", func(t *testing.T) {
		buffer.Reset()

		Debug("Test Debug")
		expected := "\x1b[90m[DEBUG]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Debug\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})

	t.Run("When passing format string", func(t *testing.T) {
		buffer.Reset()

		Debug("Test %s", "Debug2")
		expected := "\x1b[90m[DEBUG]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Debug2\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})
}

func TestInfo(t *testing.T) {
	assert := assert.New(t)
	buffer := &bytes.Buffer{}
	Stdout = buffer
	color.NoColor = false
	Now = func() time.Time {
		return time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	}

	t.Run("When not passing format string", func(t *testing.T) {
		buffer.Reset()

		Info("Test Info")
		expected := "\x1b[96m[INFO]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Info\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})

	t.Run("When passing format string", func(t *testing.T) {
		buffer.Reset()

		Info("Test %s", "Info2")
		expected := "\x1b[96m[INFO]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Info2\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})
}

func TestWarn(t *testing.T) {
	assert := assert.New(t)
	buffer := &bytes.Buffer{}
	Stderr = buffer
	color.NoColor = false
	Now = func() time.Time {
		return time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	}

	t.Run("When not passing format string", func(t *testing.T) {
		buffer.Reset()

		Warn("Test Warn")
		expected := "\x1b[93m[WARN]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Warn\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})

	t.Run("When passing format string", func(t *testing.T) {
		buffer.Reset()

		Warn("Test %s", "Warn2")
		expected := "\x1b[93m[WARN]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Warn2\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})
}

func TestError(t *testing.T) {
	assert := assert.New(t)
	buffer := &bytes.Buffer{}
	Stderr = buffer
	color.NoColor = false
	Now = func() time.Time {
		return time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	}

	t.Run("When not passing format string", func(t *testing.T) {
		buffer.Reset()

		Error("Test Error")
		expected := "\x1b[91m[ERROR]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Error\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})

	t.Run("When passing format string", func(t *testing.T) {
		buffer.Reset()

		Error("Test %s", "Error2")
		expected := "\x1b[91m[ERROR]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Error2\n"
		assert.Equal(expected, string(buffer.Bytes()))
	})
}
