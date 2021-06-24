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
