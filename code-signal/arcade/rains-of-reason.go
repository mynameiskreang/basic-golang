package main

import "strconv"

// 25: arrayReplace https://app.codesignal.com/arcade/intro/level-6/mCkmbxdMsMTjBc3Bm
func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for i, a := range inputArray {
		if a == elemToReplace {
			inputArray[i] = substitutionElem
		}
	}
	return inputArray
}

// 26: evenDigitsOnly https://app.codesignal.com/arcade/intro/level-6/6cmcmszJQr6GQzRwW
func evenDigitsOnly(n int) bool {
	mapOdd := map[string]bool{
		"1": false,
		"3": false,
		"5": false,
		"7": false,
		"9": false,
	}
	nStr := strconv.Itoa(n)
	for _, v := range nStr {
		if b, exist := mapOdd[string(v)]; exist {
			return b
		}
	}
	return true
}

//  27: variableName https://app.codesignal.com/arcade/intro/level-6/6Wv4WsrsMJ8Y2Fwno
func variableName(name string) bool {
	// a-z -> 97-122
	// A-Z -> 65-90
	// 0-9 -> 48-57
	// 95 -> _

	if name[0] >= 48 && name[0] <= 57 {
		return false
	}

	for _, s := range name {
		if s >= 0 && s <= 47 {
			return false
		}
		if s >= 58 && s <= 64 {
			return false
		}
		if s >= 91 && s <= 94 || s == 96 {
			return false
		}
		if s >= 123 && s <= 127 {
			return false
		}
	}
	return true
}

// 28: alphabeticShift https://app.codesignal.com/arcade/intro/level-6/PWLT8GBrv9xXy4Dui
func alphabeticShift(inputString string) string {
	bs := []byte(inputString)
	for i, c := range bs {
		if c+1 == 123 {
			bs[i] = 'a'
		} else {
			bs[i] = c + 1
		}
	}
	return string(bs)
}
