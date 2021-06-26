package main

// 25: arrayReplace https://app.codesignal.com/arcade/intro/level-6/mCkmbxdMsMTjBc3Bm
func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for i, a := range inputArray {
		if a == elemToReplace {
			inputArray[i] = substitutionElem
		}
	}
	return inputArray
}
