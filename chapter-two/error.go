package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// lookup at /usr/local/go/src/errors/errors.go
func (e *MyError) Error() string {
	return fmt.Sprintf("At %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"Override Error",
	}
}

func walk() error {
	return errors.New("Default Error")
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	if err := walk(); err != nil {
		fmt.Println(err)
	}
}
