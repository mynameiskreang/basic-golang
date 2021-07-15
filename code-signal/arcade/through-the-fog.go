package main

import "math"

// 30: circleOfNumbers https://app.codesignal.com/arcade/intro/level-7/vExYvcGnFsEYSt8nQ
func circleOfNumbers(n int, firstNumber int) int {
	s1 := make([]int, n/2)
	s2 := make([]int, n/2)
	for i := 0; i <= n/2-1; i++ {
		s1[i], s2[i] = i, n/2+i
	}
	for i := 0; i < len(s1); i++ {
		if firstNumber == s1[i] {
			return s2[i]
		}
		if firstNumber == s2[i] {
			return s1[i]
		}
	}
	return 0
	// best solution -> return (firstNumber + n / 2) % n
}

// 31: depositProfit https://app.codesignal.com/arcade/intro/level-7/8PxjMSncp9ApA4DAb
func depositProfit(deposit int, rate int, threshold int) int {
	fD := float64(deposit)
	fR := float64(rate)
	fT := float64(threshold)
	s := 0.0
	for i := 1; true; i++ {
		n := (fD + s) * fR / 100
		s += n
		if fD+s >= fT {
			return i
		}
	}
	return 0
	// return int(math.Ceil(math.Log(float64(threshold) / float64(deposit)) / math.Log(1.0 + float64(rate) / 100)))
}

// 32: absoluteValuesSumMinimization https://app.codesignal.com/arcade/intro/level-7/ZFnQkq9RmMiyE6qtq
func absoluteValuesSumMinimization(a []int) int {
	m := math.MaxInt32
	ans := a[0]
	for _, v := range a {
		n := getAbsoluteValuesSumMinimization(v, a)
		if n == m && v < ans {
			ans = v
		}
		if n < m {
			m = n
			ans = v
		}

	}
	return ans
	// best solution -> return A[(len(A)-1)/2]
}
func getAbsoluteValuesSumMinimization(i int, a []int) int {
	s := 0.0
	for _, v := range a {
		s += math.Abs(float64(v - i))
	}
	return int(s)
}

// 33：stringsRearrangement https://app.codesignal.com/arcade/intro/level-7/PTWhv2oWqd6p4AHB9/
func stringsRearrangement(inputArray []string) bool {
	allPerm := make([][]string, getFactorial(len(inputArray)))
	l := 0
	Perm(inputArray, func(a []string) {
		allPerm[l] = make([]string, len(inputArray))
		copy(allPerm[l], a)
		l++
	})
	for i := 0; i < len(allPerm); i++ {
		r := 0
		flag := true
		cPerm := allPerm[i]
		// fmt.Println(cPerm)
		for j := 1; j < len(cPerm); j++ {
			// fmt.Println(r,j)
			isPass := checkChar(cPerm[r], cPerm[j])
			if isPass {
				r = j
			} else {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}

	return false
}

func getFactorial(n int) int {
	sum := 1
	for i := 1; i < n+1; i++ {
		sum = sum * i
	}
	return sum
}

func checkChar(s1, s2 string) bool {
	missCount := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			missCount++
		}
	}
	if missCount >= 2 {
		return false
	}
	if missCount == 0 {
		return false
	}
	return true
}

// https://yourbasic.org/golang/generate-permutation-slice-string/
// Perm calls f with each permutation of a.
func Perm(a []string, f func([]string)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
