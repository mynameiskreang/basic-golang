package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	animals := []string{"CAT", "DOG", "FISH", "SNAKE"}
	wg := sync.WaitGroup{}
	wg.Add(len(animals))
	for _, animal := range animals {
		go printAnimal4(animal, &wg)
	}
	wg.Wait()
	endTime := time.Now()
	diffTIme := endTime.Sub(startTime)
	fmt.Println(diffTIme)
}

func printAnimal4(animal string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(animal)))
	}
	fmt.Println(animal)
	wg.Done()
}
