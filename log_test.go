package main

import (
	"bytes"
	"github.com/fatih/color"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintf(t *testing.T) {
	defer HookStdio()

	outbuf := &bytes.Buffer{}
	errbuf := &bytes.Buffer{}
	Stdout = outbuf
	Stderr = errbuf
	color.NoColor = false

	t.Run("When not passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Printf("Test Printf")

		expected := "Test Printf\n"
		assert.Equal(expected, outbuf.String())

		expected2 := ""
		assert.Equal(expected2, errbuf.String())
	})

	t.Run("When passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Printf("Test %s", "Printf2")

		expected := "Test Printf2\n"
		assert.Equal(expected, outbuf.String())

		expected2 := ""
		assert.Equal(expected2, errbuf.String())
	})
}

func TestDebug(t *testing.T) {
	defer HookStdio()

	outbuf := &bytes.Buffer{}
	errbuf := &bytes.Buffer{}
	Stdout = outbuf
	Stderr = errbuf
	color.NoColor = false

	t.Run("When not passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Debug("Test Debug")

		expected := "\x1b[90m[DEBUG]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Debug\n"
		assert.Equal(expected, outbuf.String())

		expected2 := ""
		assert.Equal(expected2, errbuf.String())
	})

	t.Run("When passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Debug("Test %s", "Debug2")

		expected := "\x1b[90m[DEBUG]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Debug2\n"
		assert.Equal(expected, outbuf.String())

		expected2 := ""
		assert.Equal(expected2, errbuf.String())
	})
}

func TestInfo(t *testing.T) {
	defer HookStdio()

	outbuf := &bytes.Buffer{}
	errbuf := &bytes.Buffer{}
	Stdout = outbuf
	Stderr = errbuf
	color.NoColor = false

	t.Run("When not passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Info("Test Info")

		expected := "\x1b[96m[INFO]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Info\n"
		assert.Equal(expected, outbuf.String())

		expected2 := ""
		assert.Equal(expected2, errbuf.String())
	})

	t.Run("When passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Info("Test %s", "Info2")

		expected := "\x1b[96m[INFO]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Info2\n"
		assert.Equal(expected, outbuf.String())

		expected2 := ""
		assert.Equal(expected2, errbuf.String())
	})
}

func TestWarn(t *testing.T) {
	defer HookStdio()

	outbuf := &bytes.Buffer{}
	errbuf := &bytes.Buffer{}
	Stdout = outbuf
	Stderr = errbuf
	color.NoColor = false

	t.Run("When not passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Warn("Test Warn")

		expected := ""
		assert.Equal(expected, outbuf.String())

		expected2 := "\x1b[93m[WARN]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Warn\n"
		assert.Equal(expected2, errbuf.String())
	})

	t.Run("When passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Warn("Test %s", "Warn2")

		expected := ""
		assert.Equal(expected, outbuf.String())

		expected2 := "\x1b[93m[WARN]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Warn2\n"
		assert.Equal(expected2, errbuf.String())
	})
}

func TestError(t *testing.T) {
	defer HookStdio()

	outbuf := &bytes.Buffer{}
	errbuf := &bytes.Buffer{}
	Stdout = outbuf
	Stderr = errbuf
	color.NoColor = false

	t.Run("When not passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Error("Test Error")

		expected := ""
		assert.Equal(expected, outbuf.String())

		expected2 := "\x1b[91m[ERROR]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Error\n"
		assert.Equal(expected2, errbuf.String())
	})

	t.Run("When passing format string", func(t *testing.T) {
		assert := assert2.New(t)

		outbuf.Reset()
		errbuf.Reset()

		Error("Test %s", "Error2")

		expected := ""
		assert.Equal(expected, outbuf.String())

		expected2 := "\x1b[91m[ERROR]\x1b[0m\x1b[97m[15:04:05] \x1b[0mTest Error2\n"
		assert.Equal(expected2, errbuf.String())
	})
}
