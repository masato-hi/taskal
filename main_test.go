package main

import (
	"bytes"
	"github.com/fatih/color"
	"os"
	"testing"
	"time"
)

var iobuffer = &bytes.Buffer{}

func TestMain(m *testing.M) {
	defer tearDown()

	setup()

	code := m.Run()

	os.Exit(code)
}

func HookStdio() {
	iobuffer.Reset()
	Stdout = iobuffer
	Stderr = iobuffer
	color.NoColor = true
}

func setup() {
	HookStdio()

	Now = func() time.Time {
		return time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	}
}

func tearDown() {
	iobuffer.Reset()
	Stdout = os.Stdout
	Stderr = os.Stderr
	color.NoColor = false

	Now = func() time.Time {
		return time.Now()
	}
}
