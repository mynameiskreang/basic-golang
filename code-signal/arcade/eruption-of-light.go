package main

import (
	"net"
	"strings"
)

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

//44: findEmailDomain https://app.codesignal.com/arcade/intro/level-10/TXFLopNcCNbJLQivP
func findEmailDomain(address string) string {
	lastAdd := strings.LastIndexAny(address, "@")
	return address[lastAdd+1:]
}

//45: buildPalindrome https://app.codesignal.com/arcade/intro/level-10/ppZ9zSufpjyzAsSEx
// I cannot pass this mission.
func reverse(s string) string {
	t := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}
	return string(t)
}

func buildPalindrome(st string) string {
	ts := reverse(st)
	for i := 0; i < len(st); i++ {
		matched := true
		j := i
		for ; j < len(st); j++ {
			if st[j] != ts[j-i] {
				matched = false
				break
			}
		}
		if matched {
			return st + ts[j-i:]
		}
	}
	return st + ts
}

//46: electionsWinners https://app.codesignal.com/arcade/intro/level-10/8RiRRM3yvbuAd3MNg
func electionsWinners(votes []int, k int) int {
	count := 0
	for i := 0; i < len(votes); i++ {
		ik := votes[i] + k
		flagK := true
		for j := 0; j < len(votes); j++ {
			if i == j {
				continue
			}
			if ik <= votes[j] {
				flagK = false
				break
			}
		}
		if flagK {
			count++
		}
	}
	return count
}

//47: isMAC48Address https://app.codesignal.com/arcade/intro/level-10/HJ2thsvjL25iCvvdm
func isMAC48Address(s string) bool {
	_, err := net.ParseMAC(s)
	return err == nil
}
