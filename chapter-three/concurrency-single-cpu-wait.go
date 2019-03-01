package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	startTime := time.Now()
	animals := []string{"CAT", "DOG", "FISH", "SNAKE"}
	wg := sync.WaitGroup{}
	wg.Add(len(animals))
	for _, animal := range animals {
		go printAnimal3(animal, &wg)
	}
	wg.Wait()
	endTime := time.Now()
	diffTIme := endTime.Sub(startTime)
	fmt.Println(diffTIme)
}

func printAnimal3(animal string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(animal)))
	}
	fmt.Println(animal)
	wg.Done()
}
