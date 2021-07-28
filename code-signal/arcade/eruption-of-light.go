package main

//43: isBeautifulString https://app.codesignal.com/arcade/intro/level-10/PHSQhLEw3K2CmhhXE
func isBeautifulString(inputString string) bool {
	mapChar := map[rune]int{}
	for _, v := range inputString {
		mapChar[v] += 1
	}
	for i := 'a'; i < 'z'; i++ {
		if mapChar[i] < mapChar[i+1] {
			return false
		}
	}
	return true
}
