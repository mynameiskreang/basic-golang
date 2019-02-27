package main

import "fmt"

func main() {
	i, j := 5, 65524
	fmt.Printf("i: %p, j: %p\n", &i, &j)
	p := &i                // point to i
	fmt.Println("p: ", *p) // read i through the pointer
	fmt.Printf("i: %p, j: %p, p: %p\n", &i, &j, p)
	*p = 21               // set i through the pointer
	fmt.Println("i: ", i) // see the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Printf("i: %p, j: %p, p: %p\n", &i, &j, p)
	fmt.Println("j: ", j) // see the new value of j
}
