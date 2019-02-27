package main

import "fmt"

func main() {
	// String
	primes := [10]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, len(primes), cap(primes))
	// Slice
	var s []int = primes[2:4] // [startIndex:targetIndex+1]
	s[0] = s[0] * 10
	fmt.Println(s, len(s), cap(s))
	fmt.Println(primes, len(primes), cap(primes))
}
