package main

import (
	"math"
)

// 19: areEquallyStrong https://app.codesignal.com/arcade/intro/level-5/g6dc9KJyxmFjB98dL
func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	if yourLeft+yourRight != friendsLeft+friendsRight {
		return false
	}

	mapStrong := make(map[int]bool)
	mapStrong[yourLeft] = true
	mapStrong[yourRight] = true
	mapStrong[friendsLeft] = true
	mapStrong[friendsRight] = true
	if len(mapStrong) > 2 {
		return false
	}

	return true
}

// 20: arrayMaximalAdjacentDifference https://app.codesignal.com/arcade/intro/level-5/EEJxjQ7oo7C5wAGjE
func arrayMaximalAdjacentDifference(inputArray []int) int {
	return arrayMaximalAdjacentDifferenceMax(inputArray, 0)
}

func arrayMaximalAdjacentDifferenceMax(inputArray []int, max int) int {
	if len(inputArray) == 1 {
		return max
	}
	result := int(math.Abs(float64(inputArray[0] - inputArray[1])))
	if result > max {
		max = result
	}
	return arrayMaximalAdjacentDifferenceMax(inputArray[1:], max)
}
