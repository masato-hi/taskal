package main

import (
	"strconv"
	"strings"
	"unicode"
)

func TrimTailingSpace(str string) string {
	return strings.TrimRightFunc(str, unicode.IsSpace)
}

func QuoteString(str string) string {
	return strconv.Quote(str)
}
