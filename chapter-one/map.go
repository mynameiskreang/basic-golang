package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	value1, isExsit1 := m["Answer"]
	value2, isExsit2 := m["Question"]
	fmt.Println("The value:", value1, "Present?", isExsit1)
	fmt.Println("The value:", value2, "Present?", isExsit2)

	delete(m, "Answer")
	value1, isExsit1 = m["Answer"]
	fmt.Println("The value:", value1, "Present?", isExsit1)

}
