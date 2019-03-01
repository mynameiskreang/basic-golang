package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	startTime := time.Now()
	animals := []string{"CAT", "DOG", "FISH", "SNAKE"}
	for _, animal := range animals {
		go printAnimal2(animal)
	}
	endTime := time.Now()
	diffTIme := endTime.Sub(startTime)
	fmt.Println(diffTIme)
}

func printAnimal2(animal string) {
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(animal)))
	}
	fmt.Println(animal)
}
