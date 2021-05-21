package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
)

type Age struct {
}

func (r Age) String() string {
	return "Now, you are a man."
}

func (r Age) Set(s string) error {
	age, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if age < 18 {
		return errors.New("You are young.")
	}
	return nil
}

// Step 1:
// เป็นการเขียน custom flag โดยไปลอก Value
// type Value interface {
//	String() string
//	Set(string) error
// }
// ดังนั้น struct age ต้อง implement String() และ Set()
func main() {
	age := Age{}
	flag.Var(&age, "age", "help message")
	flag.Parse()
	fmt.Println(age)
}
