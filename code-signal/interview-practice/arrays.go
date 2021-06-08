package main

// Arrays: firstDuplicate https://app.codesignal.com/interview-practice/task/pMvymcahZ8dY4g75q/solutions
func firstDuplicate(a []int) int {
	set := make(map[int]bool)
	for _, value := range a {
		if set[value] {
			return value
		} else {
			set[value] = true
		}
	}
	return -1
}
