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

// 29: chessBoardCellColor https://app.codesignal.com/arcade/intro/level-6/t97bpjfrMDZH8GJhi
func chessBoardCellColor(cell1 string, cell2 string) bool {
	// a c e g 1 3 5 7 -> brown
	// b d f h 2 4 6 8 -> brown
	// 65 67 69 71, 49 51 53 55
	// 66 68 70 72,
	c1Yellow := true
	c2Yellow := true
	bcell1 := []byte(cell1)
	bcell2 := []byte(cell2)
	if (bcell1[0]%2 == 1 && bcell1[1]%2 == 1) || (bcell1[0]%2 == 0 && bcell1[1]%2 == 0) {
		// brown
		c1Yellow = false
	}
	if (bcell2[0]%2 == 1 && bcell2[1]%2 == 1) || (bcell2[0]%2 == 0 && bcell2[1]%2 == 0) {
		// brown
		c2Yellow = false
	}
	return c1Yellow == c2Yellow
}
