package main

import (
	"net/url"
	"strings"
)

type stringReader interface {
	ReadString(byte) (string, error)
}

func ReadInput(r stringReader) (string, error) {
	// text, err := r.ReadString('^')
	return r.ReadString('^')
}

func FormatString(s string) string {
	if len(s) > 120 {
		s = s[0:120]
	}
	tabNewLineRemoval := strings.Replace(s, "\n\t", " ", -1)
	newLineRmove := strings.Replace(tabNewLineRemoval, "\n\t", " ", -1)
	queryEscaped := url.QueryEscape(newLineRmove)
	return strings.Replace(queryEscaped, "\t", "", -1)
}
