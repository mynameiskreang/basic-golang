package main

import (
	"fmt"
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
