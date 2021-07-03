package main

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
