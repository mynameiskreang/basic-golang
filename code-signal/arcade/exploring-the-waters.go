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

// 15: addBorder https://app.codesignal.com/arcade/intro/level-4/ZCD7NQnED724bJtjN
func addBorder(picture []string) []string {
	HL := len(picture[0]) + 2
	output := []string{getStar(HL)}
	for i := 0; i < len(picture); i++ {
		picture[i] = "*" + picture[i] + "*"
	}
	output = append(output, picture...)
	output = append(output, getStar(HL))
	return output
}
func getStar(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s = s + "*"
	}
	return s
}
