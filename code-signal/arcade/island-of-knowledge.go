package main

import (
	"math"
	"strconv"
	"strings"
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

// 21: isIPv4Address https://app.codesignal.com/arcade/intro/level-5/veW5xJednTy4qcjso
func isIPv4Address(inputString string) bool {

	// Interesting Solution
	// return net.ParseIP(inputString) != nil
	// Just using package "net"

	ips := strings.Split(inputString, ".")
	if len(ips) != 4 {
		return false
	}
	for _, ip := range ips {
		ipN, err := strconv.Atoi(ip)
		if err != nil {
			return false
		}
		if ipN < 0 || ipN > 255 || len(ip) == 0 || len(ip) != len(strconv.Itoa(ipN)) {
			return false
		}
	}
	return true
}

// 22: avoidObstacles https://app.codesignal.com/arcade/intro/level-5/XC9Q2DhRRKQrfLhb5
func avoidObstacles(inputArray []int) int {
	start := 2
	for i := 0; i < len(inputArray); {
		stand := inputArray[i]
		if stand%start == 0 || stand == start {
			start++
			i = 0
			continue
		} else {
			i++
		}
	}
	return start
}

// 23:  boxBlur https://app.codesignal.com/arcade/intro/level-5/5xPitc3yT3dqS7XkP
func boxBlur(image [][]int) [][]int {
	maxX := len(image[0])
	maxY := len(image)
	// fmt.Println(maxX,maxY)
	// fmt.Println()
	output := [][]int{}
	// yLine:=[]int{}
	for y := 0; y < maxY-2; y++ {
		xLine := []int{}
		for x := 0; x < maxX-2; x++ {
			sum := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					// fmt.Printf(" %d ",image[i+y][j+x])
					sum = sum + image[i+y][j+x]
				}
				// fmt.Println()
			}
			result := sum / 9
			xLine = append(xLine, result)
			// fmt.Println("result", result)
		}
		output = append(output, xLine)
	}
	return output
}

// มี solution ที่น่าสนใจก็คือ
func boxBlurSolution(image [][]int) [][]int {
	// Step 1 prepare array
	b := make([][]int, len(image)-2)
	for i := range b {
		b[i] = make([]int, len(image[0])-2)
	}
	// Step 2 loop array b & summary
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			b[i][j] = (image[i][j] + image[i+1][j] + image[i+2][j] +
				image[i][j+1] + image[i+1][j+1] + image[i+2][j+1] +
				image[i][j+2] + image[i+1][j+2] + image[i+2][j+2]) / 9
		}
	}
	return b
}

// 24: minesweeper https://app.codesignal.com/arcade/intro/level-5/ZMR5n7vJbexnLrgaM
func minesweeper(matrix [][]bool) [][]int {
	output:=make([][]int,len(matrix))
	for i:= range matrix{
		output[i] = make([]int,len(matrix[0]))
	}
	// -1,-1  0,-1  +1,-1
	// -1, 0  j, i  +1, 0
	// -1,+1  0,+1  +1,+1
	for i:=0;i<len(output);i++{
		for j:=0;j<len(output[0]);j++{
			count:=0
			// Line1
			// -1,-1
			if i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1]{
				count++
			}
			// 0,-1
			if j-1 >= 0 && matrix[i][j-1]{
				count++
			}
			// +1,-1
			if i+1 < len(matrix) && j-1 >= 0 && matrix[i+1][j-1]{
				count++
			}
			// Line2
			// -1,0
			if i-1 >= 0 && matrix[i-1][j]{
				count++
			}
			// +1,0
			if i+1 < len(matrix)  && matrix[i+1][j]{
				count++
			}
			// Line3
			// -1,+1
			if i-1 >= 0 && j+1 < len(matrix[0]) && matrix[i-1][j+1]{
				count++
			}
			// 0,+1
			if j+1 < len(matrix[0]) && matrix[i][j+1]{
				count++
			}
			// +1,+1
			if i+1 < len(matrix) && j+1 < len(matrix[0]) && matrix[i+1][j+1]{
				count++
			}
			output[i][j] = count
		}
	}
	return output
}

// solution นี้น่าสนใจมาก
// มีการประกาศ neighbors เพื่อทำ mapping ไว้ก่อน
// มันทำให้ code อ่านง่าย
// https://app.codesignal.com/arcade/intro/level-5/ZMR5n7vJbexnLrgaM/solutions?solutionId=L65AiqxdvN5fyTrnR
{
	var neighbors = [][]int{
		{ -1, -1}, {0, -1}, {1, -1},
		{ -1,  0},          {1, 0},
		{ -1,  1}, {0,  1}, {1, 1},
	}

	func mine(field [][]bool, x, y int) int {
		if (x < 0 || x >= len(field[0])) {
			return 0
		}
		if (y < 0 || y >= len(field)) {
			return 0
		}
		if (field[y][x]) {
			return 1
		}
		return 0
	}

	func eval(field [][]bool, x, y int) int {
		sum := 0
		for _, n := range neighbors {
			sum += mine(field, x + n[0], y + n[1])
		}
		return sum
	}

	func minesweeperL65AiqxdvN5fyTrnR(field [][]bool) [][]int {
		var ret [][]int
		for y := range field {
			row := make([]int, len(field[0]))
			for x := range row {
				row[x] = eval(field, x, y)
			}
			ret = append(ret, row)
		}
		return ret
	}
}
