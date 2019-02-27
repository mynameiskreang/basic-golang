package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	default:
		fmt.Println("Good evening.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() < 12:
		fmt.Println("Good morning!")

	}
}
