package main

// 14: alternatingSums https://app.codesignal.com/arcade/intro/level-4/cC5QuL9fqvZjXJsW9
func alternatingSums(a []int) []int {
	if len(a)%2 != 0 {
		a = append(a, 0)
	}
	return forkAlternatingSums(a)
}
func forkAlternatingSums(a []int) []int {
	if len(a) <= 2 {
		return a
	}
	a[2] = a[0] + a[2]
	return forkAlternatingSums(a[1:])
}
