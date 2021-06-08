package main

import (
	"fmt"
	"strconv"
)

// 1: add https://app.codesignal.com/arcade/intro/level-1/jwr339Kq6e3LQTsfa
func add(param1 int, param2 int) int {
	return param1 + param2
}

// 2: centuryFromYear https://app.codesignal.com/arcade/intro/level-1/egbueTZRRL5Mm4TXN
func centuryFromYear(year int) int {
	yearStr := strconv.Itoa(year)
	for len(yearStr) < 4 {
		yearStr = "0" + yearStr
	}
	fmt.Println(yearStr[:2], yearStr[2:])
	f, _ := strconv.Atoi(yearStr[:2])
	s, _ := strconv.Atoi(yearStr[2:])
	if s > 0 {
		return f + 1
	}
	return f
}

// 3: checkPalindrome https://app.codesignal.com/arcade/intro/level-1/s5PbmwxfECC52PWyQ
func checkPalindrome(inputString string) bool {
	if len(inputString) == 1 {
		return true
	}
	if len(inputString)%2 == 0 {
		cut := len(inputString) / 2
		f := inputString[:cut]
		s := inputString[cut:]
		sr := Reverse(s)
		fmt.Println(cut, f, s, sr)
		if f == sr {
			return true
		}
	} else {
		cut := len(inputString) / 2
		f := inputString[:cut]
		s := inputString[cut+1:]
		sr := Reverse(s)
		fmt.Println(cut, f, s, sr)
		if f == sr {
			return true
		}
	}
	return false
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//
