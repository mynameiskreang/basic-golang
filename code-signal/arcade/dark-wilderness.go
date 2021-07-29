package main

import (
	"math"
	"strconv"
	"strings"
)

//38：growingPlant https://app.codesignal.com/arcade/intro/level-9/xHvruDnQCx7mYom3T
func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	if upSpeed >= desiredHeight {
		return 1
	}
	output := ((desiredHeight - upSpeed) / (upSpeed - downSpeed))
	if output == 0 {
		output++
	}
	return output + 1
}

//39: knapsackLight https://app.codesignal.com/arcade/intro/level-9/r9azLYp2BDZPyzaG2
func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	// Step1 order
	orderValue := make([]*int, 2)
	orderWeight := make([]*int, 2)
	if value1 > value2 {
		orderValue[0] = &value1
		orderValue[1] = &value2
		orderWeight[0] = &weight1
		orderWeight[1] = &weight2
	} else {
		orderValue[0] = &value2
		orderValue[1] = &value1
		orderWeight[0] = &weight2
		orderWeight[1] = &weight1
	}
	// Step2 calculate
	sum := 0
	for i, _ := range orderValue {
		if maxW-*orderWeight[i] >= 0 {
			maxW = maxW - *orderWeight[i]
			sum += *orderValue[i]
		}
	}
	return sum
}

//40: longestDigitsPrefix https://app.codesignal.com/arcade/intro/level-9/AACpNbZANCkhHWNs3
func longestDigitsPrefix(inputString string) string {
	output := ""
	for _, v := range inputString {
		if v >= 48 && v <= 57 {
			output += string(v)
		} else {
			break
		}
	}
	return output
}

//41: digitDegree https://app.codesignal.com/arcade/intro/level-9/hoLtYWbjdrD2PF6yo
func digitDegree(n int) int {
	ns := strconv.Itoa(n)
	count := 0
	for count = 0; ; count++ {
		if len(ns) == 1 {
			break
		}
		ns = sumString(ns)

	}
	return count
}

func sumString(input string) string {
	sum := 0
	for _, v := range input {
		val, _ := strconv.Atoi(string(v))
		sum += val
	}
	return strconv.Itoa(sum)
}

//42: bishopAndPawn https://app.codesignal.com/arcade/intro/level-9/6M57rMTFB9MeDeSWo
func bishopAndPawn(bishop string, pawn string) bool {
	mapChar := make(map[byte]int, 8)
	mapChar['a'] = 1
	mapChar['b'] = 2
	mapChar['c'] = 3
	mapChar['d'] = 4
	mapChar['e'] = 5
	mapChar['f'] = 6
	mapChar['g'] = 7
	mapChar['h'] = 8
	b1 := mapChar[bishop[0]]
	b2, _ := strconv.Atoi(string(bishop[1]))
	p1 := mapChar[pawn[0]]
	p2, _ := strconv.Atoi(string(pawn[1]))
	if math.Abs(float64(b1)-float64(b2)) == math.Abs(float64(p1)-float64(p2)) {
		return true
	}
	if b1+b2 == p1+p2 {
		return true
	}
	return false
}

//43：findEmailDomain https://app.codesignal.com/arcade/intro/level-10/TXFLopNcCNbJLQivP
func findEmailDomain(address string) string {
	lastAdd := strings.LastIndexAny(address, "@")
	return address[lastAdd+1:]
}
