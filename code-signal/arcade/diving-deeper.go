package main

//34: extractEachKth https://app.codesignal.com/arcade/intro/level-8/3AgqcKrxbwFhd3Z3R/
func extractEachKth(inputArray []int, k int) []int {
	if k > len(inputArray) {
		return inputArray
	}
	outputArray := []int{}
	for i := 0; i < len(inputArray); i++ {
		if (i+1)%k != 0 {
			outputArray = append(outputArray, inputArray[i])
		}
	}
	return outputArray
}
