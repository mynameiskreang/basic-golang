package main

import (
	"fmt"
	"unicode"
)

//48: isDigit https://app.codesignal.com/arcade/intro/level-11/Zr2XXEpkBPtYWqPJu
func isDigit(symbol string) bool {
	return unicode.IsDigit(rune(symbol[0]))
}

//49: lineEncoding https://app.codesignal.com/arcade/intro/level-11/o2uq6eTuvk7atGadR
func lineEncoding(s string) string {
	checker := rune(s[0])
	result := ""
	count := 0
	for _, r := range s {
		if checker == r {
			count++
			continue
		} else {
			if count > 1 {
				result = result + fmt.Sprintf("%d%c", count, checker)
			} else {
				result = result + fmt.Sprintf("%c", checker)
			}
			count = 1
			checker = r
		}
	}
	if count > 1 {
		result = result + fmt.Sprintf("%d%c", count, checker)
	} else {
		result = result + fmt.Sprintf("%c", checker)
	}
	return result
}
