package main

import "fmt"

// 4: adjacentElementsProduct https://app.codesignal.com/arcade/intro/level-2/xzKiBHjhoinnpdh6m
func adjacentElementsProduct(inputArray []int) int {
	s := inputArray[0] * inputArray[1]
	if len(inputArray) == 2 {
		return s
	}
	for i := 0; i < len(inputArray)-1; i++ {
		if s < inputArray[i]*inputArray[i+1] {
			s = inputArray[i] * inputArray[i+1]
		}
	}
	return s
}

// 5: shapeArea https://app.codesignal.com/arcade/intro/level-2/yuGuHvcCaFCKk56rJ
func shapeArea(n int) int {
	result := 0
	for i := n; i > 0; i-- {
		if i == n {
			result = result + i*2 - 1
		} else {
			result = result + (i*2-1)*2
		}
	}
	return result
}

// 6: makeArrayConsecutive2 https://app.codesignal.com/arcade/intro/level-2/bq2XnSr5kbHqpHGJC
func makeArrayConsecutive2(statues []int) int {
	master := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	max := findMax(statues, master)
	min := findMin(statues, master)
	result := max - min - len(statues) + 1
	fmt.Println(max, min, result)
	// [6, 2, 3, 8] => 3 [2, 3, '4', '5', 6, 7, 8]
	// [0, 3] => 2 [0, 1, 2, 3]
	// [5, 4, 6] => 0
	// [6, 4] => 2
	// [1] => 0
	return result
}
func findMax(statues, master []int) int {
	max := 0
	for _, m := range master {
		for _, s := range statues {
			if m == s && s > max {
				max = s
			}
		}
	}
	return max
}
func findMin(statues, master []int) int {
	min := 20
	for _, m := range master {
		for _, s := range statues {
			if m == s && s < min {
				min = s
			}
		}
	}
	return min
}

// 7: almostIncreasingSequence https://app.codesignal.com/arcade/intro/level-2/2mxbGwLzvkTCKAJMG
func almostIncreasingSequence(sequence []int) bool {
	var x int = 0
	var y int = 0

	for i := 1; i < len(sequence)-1; i++ {
		if sequence[i] <= sequence[i-1] {
			x++
		}
		if sequence[i+1] <= sequence[i-1] {
			y++
		}
	}

	if sequence[len(sequence)-1] <= sequence[len(sequence)-2] {
		x++
	}

	if x <= 1 && y <= 1 {
		return true
	} else {
		return false
	}
}

// 8: matrixElementsSum https://app.codesignal.com/arcade/intro/level-2/xskq4ZxLyqQMCLshr
func matrixElementsSum(matrix [][]int) int {
	xAxis := len(matrix[0])
	sumRooms := 0
	for j := 0; j < xAxis; j++ {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] == 0 {
				break
			}
			// fmt.Println(matrix[i][j])
			sumRooms += matrix[i][j]
		}
	}
	return sumRooms
}
