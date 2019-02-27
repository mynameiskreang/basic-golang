package main

import (
	"fmt"
	"math"
)

func main() {
	pow := make([]int, 10)
	for index := range pow {
		pow[index] = int(math.Pow10(index))
	}

	for index, value := range pow {
		fmt.Printf("10^%d = %d\n", index, value)
	}
}
