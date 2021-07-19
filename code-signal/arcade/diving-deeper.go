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

//35: firstDigit https://app.codesignal.com/arcade/intro/level-8/rRGGbTtwZe2mA8Wov
func firstDigit(inputString string) string {
	mapChar := map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
	for _, v := range inputString {
		if mapChar[v] {
			return string(v)
		}
	}
	return ""
}

//36: differentSymbolsNaive https://app.codesignal.com/arcade/intro/level-8/8N7p3MqzGQg5vFJfZ
func differentSymbolsNaive(s string) int {
	mapSymbols := map[rune]bool{}
	for _, v := range s {
		mapSymbols[v] = true
	}
	return len(mapSymbols)
}

//37: arrayMaxConsecutiveSum https://app.codesignal.com/arcade/intro/level-8/Rqvw3daffNE7sT7d5
func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	max, sum := inputArray[0], 0
	for i := 0; i < len(inputArray); i++ {
		if i < k {
			sum += inputArray[i]
		} else {
			sum -= inputArray[i-k]
			sum += inputArray[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
