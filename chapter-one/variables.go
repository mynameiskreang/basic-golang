package main

import (
	"basic-golang/chapter-one/mylib"
	"fmt"
)

func main() {
	fmt.Println("mylib.X", mylib.X)
	sum := mylib.Sum(3)
	fmt.Println("sum", sum)

	mylib.X = mylib.X * 10
	sum = mylib.Sum(3)
	fmt.Println("sum", sum)
}
