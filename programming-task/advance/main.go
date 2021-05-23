package main

import (
	"fmt"
	"strings"
)

// https://programming.in.th/task/rev2_problem.php?pid=2007
// input
// 4 11
// cAda
// AbrAcadAbRa
// output
// 2
func main() {
	var g, s int
	fmt.Scanf("%d %d", &g, &s)
	W := make([]byte, g)
	S := make([]byte, s)

	for i := 0; i < g; i++ {
		fmt.Scanf("%c", &W[i])
	}

	var newLine byte
	fmt.Scanf("%c", &newLine)

	for j := 0; j < s; j++ {
		fmt.Scanf("%c", &S[j])
	}

	fmt.Println("W", len(W), cap(W), "S", len(S), cap(S), string(W), string(S))

	fact := permutation(len(W), W)
	fact = append(fact, string(W))
	fmt.Println(len(fact), fact)

	count := 0
	for _, s2 := range fact {
		if strings.Contains(string(S), s2) {
			count++
		}
	}
	fmt.Println(count)
}

func Fact(n int) int {
	if n == 1 {
		return n
	}
	return n * Fact(n-1)
}

var temp = 1
var facts []string

func permutation(k int, A []byte) []string {
	//facts = append(facts, string(A))
	if k == 1 {
		//facts = append(facts, string(A))
		return facts
	}
	permutation(k-1, A)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			A[i], A[k-1] = A[k-1], A[i]
		} else {
			A[0], A[k-1] = A[k-1], A[0]
		}
		fmt.Println(temp, string(A))
		temp++
		facts = append(facts, string(A))
		permutation(k-1, A)
	}
	return facts
}
