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
