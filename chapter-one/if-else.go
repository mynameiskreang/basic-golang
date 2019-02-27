package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// ไม่สามารถใช้ตัวแปร v ได้ เพราะอยู่คนละ scope
	return lim
}

func main() {
	fmt.Println(sqrt(10.0))
	fmt.Println(pow(11.0, 20.0, 10.0))
}
