package main

import (
	"fmt"
	"strings"
)

// 14: alternatingSums https://app.codesignal.com/arcade/intro/level-4/cC5QuL9fqvZjXJsW9
func alternatingSums(a []int) []int {
	if len(a)%2 != 0 {
		a = append(a, 0)
	}
	return forkAlternatingSums(a)
}
func forkAlternatingSums(a []int) []int {
	if len(a) <= 2 {
		return a
	}
	a[2] = a[0] + a[2]
	return forkAlternatingSums(a[1:])
}

// 15: addBorder https://app.codesignal.com/arcade/intro/level-4/ZCD7NQnED724bJtjN
func addBorder(picture []string) []string {
	HL := len(picture[0]) + 2
	output := []string{getStar(HL)}
	for i := 0; i < len(picture); i++ {
		picture[i] = "*" + picture[i] + "*"
	}
	output = append(output, picture...)
	output = append(output, getStar(HL))
	return output
}
func getStar(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s = s + "*"
	}
	return s
}

// 16: areSimilar https://app.codesignal.com/arcade/intro/level-4/xYXfzQmnhBvEKJwXP
func areSimilar(a []int, b []int) bool {
	// check member
	memberA := countSameNumber(a)
	memberB := countSameNumber(b)
	fmt.Println(memberA, memberB)
	for _, v := range a {
		if memberA[v] != memberB[v] {
			return false
		}
	}

	// check swapping
	countMiss := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			countMiss++
		}
	}
	if countMiss >= 3 {
		return false
	}
	return true
}

func countSameNumber(a []int) map[int]int {
	mapAB := make(map[int]int)
	for _, v := range a {
		if val, ok := mapAB[v]; ok {
			mapAB[v] = val + 1
		} else {
			mapAB[v] = 1
		}
	}
	return mapAB
}

// 17: arrayChange https://app.codesignal.com/arcade/intro/level-4/xvkRbxYkdHdHNCKjg
func arrayChange(inputArray []int) int {
	step := 0
	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] <= inputArray[i-1] {
			gap := inputArray[i-1] - inputArray[i] + 1
			inputArray[i] += gap
			step += gap
		}
	}
	return step
}

// 18: palindromeRearranging https://app.codesignal.com/arcade/intro/level-4/Xfeo7r9SBSpo3Wico
func palindromeRearranging(inputString string) bool {
	mapChar := make(map[string]bool)
	countOdd := 0
	for _, v := range inputString {
		c := strings.Count(inputString, string(v))
		if _, ok := mapChar[string(v)]; ok == false {
			mapChar[string(v)] = true
			if c%2 != 0 {
				fmt.Println(string(v), c)
				countOdd++
			}
			if countOdd > 1 {
				return false
			}
		}
	}
	return true
}
