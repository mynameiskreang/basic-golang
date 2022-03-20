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

// 55: differentSquares https://app.codesignal.com/arcade/intro/level-12/fQpfgxiY6aGiGHLtv
func differentSquares(matrix [][]int) int {
	x := 0
	y := 0
	mapStr := make(map[string]bool)
	for {
		if x+1 < len(matrix[0]) && y+1 < len(matrix) {
			n1 := matrix[y][x]
			n2 := matrix[y+1][x]
			n3 := matrix[y][x+1]
			n4 := matrix[y+1][x+1]
			str := fmt.Sprintf("%s%s%s%s", n1, n2, n3, n4)
			mapStr[str] = true
		}
		x++
		if x == len(matrix[0]) {
			y++
			x = 0
		}
		if y == len(matrix) {
			break
		}
	}
	return len(mapStr)
}

//56: digitsProduct

//57: fileNaming https://app.codesignal.com/arcade/intro/level-12/sqZ9qDTFHXBNrQeLC
func fileNaming(names []string) []string {
	output := []string{}
	count := 0
	name := ""
	for i := 0; i < len(names); {
		if count > 0 {
			name = fmt.Sprintf("%s(%d)", names[i], count)
		} else {
			name = fmt.Sprintf("%s", names[i])
		}
		if isExist(output, name) {
			count++
			continue
		}
		output = append(output, name)
		name = ""
		count = 0
		i++
	}
	return output
}

func isExist(output []string, name string) bool {
	for _, o := range output {
		if name == o {
			return true
		}
	}
	return false
}

// 58 messageFromBinaryCode https://app.codesignal.com/arcade/intro/level-12/sCpwzJCyBy2tDSxKW/solutions
func messageFromBinaryCode(code string) string {
	output := ""
	temp := ""
	for i, v := range code {
		temp += string(v)
		if i%8 == 7 {
			dec, _ := strconv.ParseInt(temp, 2, 64)
			output += fmt.Sprintf("%c", dec)
			temp = ""
		}
	}
	return output
}

// 59 spiralNumbers https://app.codesignal.com/arcade/intro/level-12/uRWu6K8E7uLi3Qtvx
func spiralNumbers(n int) [][]int {
	var min, max, i, j int
	min = 0
	i = 0
	j = 0
	max = n - 1
	output := make([][]int, n)
	for o := range output {
		output[o] = make([]int, n)
	}
	for l := 1; l <= n*n; l++ {
		output[i][j] = l
		if i == min && j < max {
			j++
			continue
		}
		if i < max && j == max {
			i++
			continue
		}
		if i == max && j > min {
			j--
			continue
		}
		if i > min && j == min {
			i--
			if output[i][j] != 0 {
				min++
				max--
				l--
				i++
			}
			continue
		}
	}
	return output
}

// 60 sudoku https://app.codesignal.com/arcade/intro/level-12/tQgasP8b62JBeirMS
func sudoku(grid [][]int) bool {
	// case row, column
	for i := 0; i < 9; i++ {
		mapI := make(map[int]bool)
		mapJ := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if mapI[grid[i][j]] || mapJ[grid[j][i]] {
				return false
			}
			mapI[grid[i][j]] = true
			mapJ[grid[j][i]] = true
		}
	}

	// case 3x3
	for x := 0; x < 9; x = x + 3 {
		for y := 0; y < 9; y = y + 3 {
			mapIJ := make(map[int]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if mapIJ[grid[i+x][j+y]] {
						return false
					}
					mapIJ[grid[i+x][j+y]] = true
				}
			}
		}
	}

	return true
}
