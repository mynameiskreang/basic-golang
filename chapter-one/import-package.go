package main

import (
	"basic-golang/chapter-one/mylib"
	"fmt"
)

func main() {
	xy := mylib.PublicAdd(10, 20)
	fmt.Println("mylib.PublicAdd", xy)

	// Cant Access privateAdd
	//xy := mylib.privateAdd(10, 20)
	//fmt.Println("mylib.privateAdd", xy)
}
