package main

import (
	"fmt"
	"strings"
)

func main() {
	board := make([][]string, 2)
	board = [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "A"
	board[1][1] = "A"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("------")
	board = append(board, []string{"_", "_", "A"})
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
