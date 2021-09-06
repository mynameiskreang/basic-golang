package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// 52: longestWord https://app.codesignal.com/arcade/intro/level-12/s9qvXv4yTaWg8g4ma
// In my opinion, the best answer is regexp (https://app.codesignal.com/arcade/intro/level-12/s9qvXv4yTaWg8g4ma/solutions?solutionId=A82WsWxcSQKpiKSX4)
//func longestWord(text string) string {
//	longest := ""
//	re := regexp.MustCompile("[A-Za-z]+")
//	for _, elem := range re.FindAllString(text, -1) {
//		if len(elem) > len(longest) {
//			longest = elem
//		}
//	}
//	return longest
//}
func longestWord(text string) string {
	tmp := ""
	ans := ""
	for _, t := range text {
		if (t >= 65 && t <= 90) || (t >= 97 && t <= 122) {
			tmp += string(t)
		} else {
			fmt.Println(tmp)
			if len(tmp) > len(ans) {
				ans = tmp
			}
			tmp = ""
		}
	}
	if len(tmp) > len(ans) {
		ans = tmp
	}
	return ans
}

// 53: validTime https://app.codesignal.com/arcade/intro/level-12/ywMyCTspqGXPWRZx5
func validTime(time string) bool {
	v1, _ := strconv.Atoi(string(time[:2]))
	v2, _ := strconv.Atoi(string(time[3:]))
	if v1 >= 24 {
		return false
	}
	if v2 > 60 {
		return false
	}
	if v1 == 0 && v2 == 60 {
		return false
	}
	return true
}

// 54: sumUpNumbers https://app.codesignal.com/arcade/intro/level-12/YqZwMJguZBY7Hz84T
func sumUpNumbers(inputString string) int {
	sum := 0
	chain := ""
	for _, i := range inputString {
		if unicode.IsDigit(i) {
			chain += string(i)
		} else {
			v, _ := strconv.Atoi(chain)
			sum += v
			chain = ""
		}
	}
	v, _ := strconv.Atoi(chain)
	sum += v
	chain = ""
	return sum
}
