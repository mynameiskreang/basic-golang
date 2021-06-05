package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{2, 100}
	set := make(map[int]bool)

	for i := 0; i < input[1]; i++ {
		if i%2 == 0 {
			continue
		}
		for j := input[0]; j <= i; j++ {
			state := j
			if i%j == 0 {
				break
			}
			if state == i-1 {
				println(i)
				set[i] = true
			}
		}
	}
	//fmt.Println(set)
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	isPrime(n + 1)
	return false
}

func GCDAndLCM() {
	input := []int{56, 84, 140}
	var max, min float64
	for i := 0; i < len(input)-1; i++ {
		max = math.Max(float64(input[i]), float64(input[i+1])) // กำหนดเป็นตัวตั้ง
		min = math.Min(float64(input[i]), float64(input[i+1])) // กำหนดเป็นตัวหาร
	}
	fmt.Println(max, min)

	gcd := findGCD(int(max), int(min))
	lcm := findLCM(int(max), int(min))
	fmt.Println(gcd, lcm)
}

// หา  ห.ร.ม.
// ตัวอย่าง   จงหา ห.ร.ม. ของ  56   84  และ 140
// วิธีทำ            56 = 2 x 2 x 7
// 84 = 2 x 2 x 3 x 7
// 140 = 2 x 2 x 5 x 7
// เลือกตัวที่ซ้ำกัน  ที่อยู่ทั้ง 56 84และ 140 ตัวทีซ้ำกันเอามาซ้ำละ 1 ตัว
// คือ  มีเลข  2   เลข  2 และ เลข 7
// ดังนั้น       ห.ร.ม.   =  2 x 2 x 7

//func horRorMor() {
//	input := []int{56, 84, 140}
//	h := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
//}

func findGCD(num1, num2 int) int {
	if num1 == 0 {
		return num2
	}
	if num2 == 0 {
		return num1
	}
	dividend := math.Max(float64(num1), float64(num2)) // กำหนดเป็นตัวตั้ง
	divisor := math.Min(float64(num1), float64(num2))  // กำหนดเป็นตัวหาร
	fraction := int(dividend) % int(divisor)
	if fraction != 0 {
		return findGCD(int(divisor), fraction)
	} else {
		return int(divisor)
	}
}

func findLCM(num1, num2 int) int {
	//56/28*84
	return num1 * num2 / findGCD(num1, num2)
}
