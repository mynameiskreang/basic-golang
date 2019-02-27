package main

import "fmt"

func main() {
	sum := 1
	for i := 0; i < 5; i++ { // For
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 { // While
		sum += sum
	}
	fmt.Println(sum)
}
