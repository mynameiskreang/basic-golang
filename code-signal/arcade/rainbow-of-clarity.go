package main

import "unicode"

//48: isDigit https://app.codesignal.com/arcade/intro/level-11/Zr2XXEpkBPtYWqPJu/import "unicode"
func isDigit(symbol string) bool {
	return unicode.IsDigit(rune(symbol[0]))
}
