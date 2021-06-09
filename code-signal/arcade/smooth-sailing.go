package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 9: allLongestStrings https://app.codesignal.com/arcade/intro/level-3/fzsCQGYbxaEcTr2bL
func allLongestStrings(inputArray []string) []string {
	mapLen := make(map[int][]string)
	mapLen[0] = []string{}
	max := 0

	for _, v := range inputArray {
		l := len(v)
		if l > max {
			max = l
		}
		mapLen[l] = append(mapLen[l], v)
	}

	return mapLen[max]
}

// 10: commonCharacterCount https://app.codesignal.com/arcade/intro/level-3/JKKuHJknZNj4YGL32
// import "sort"
type StringToRune []rune

func (s StringToRune) Len() int           { return len(s) }
func (s StringToRune) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s StringToRune) Less(i, j int) bool { return s[i] < s[j] }
func commonCharacterCount(s1 string, s2 string) int {
	// rs1:=[]rune(s1)
	// sort.Sort(StringToRune(rs1))
	// rs2:=[]rune(s2)
	// sort.Sort(StringToRune(rs2))
	count := 0
	start := 97
	end := 122
	// fmt.Println(string(rs1),string(rs2))
	// fmt.Println([]rune("a"),[]rune("z"))
	// for i:=0;i<len(s1);i++{
	// if rs1[i] == rs2[i]{
	// count++
	// }
	// }

	for i := start; i <= end; i++ {
		cs1 := strings.Count(s1, fmt.Sprintf("%c", i))
		cs2 := strings.Count(s2, fmt.Sprintf("%c", i))
		if cs1 != 0 && cs2 != 0 {
			if cs1 > cs2 {
				count = count + cs2
			} else {
				count = count + cs1
			}
		}
	}

	return count
}

// 11: isLucky https://app.codesignal.com/arcade/intro/level-3/3AdBC97QNuhF6RwsQ
func isLucky(n int) bool {
	nStr := strconv.Itoa(n)
	fmt.Println(nStr)
	halfLen := len(nStr) / 2
	sumL := 0
	sumR := 0
	for i := 0; i < len(nStr[:halfLen]); i++ {
		// fmt.Println(string(nStr[i]))
		l, _ := strconv.Atoi(string(nStr[i]))
		sumL += l
	}
	// fmt.Println("----------")
	for i := halfLen; i < len(nStr); i++ {
		// fmt.Println(string(nStr[i]))
		r, _ := strconv.Atoi(string(nStr[i]))
		sumR += r
	}
	return sumL == sumR
}

// 12: sortByHeight https://app.codesignal.com/arcade/intro/level-3/D6qmdBL2NYz49XHwM
func sortByHeight(a []int) []int {
	minusOne := []int{}
	normalNumber := []int{}
	// Split Data
	for i := 0; i < len(a); i++ {
		if a[i] != -1 {
			normalNumber = append(normalNumber, a[i])
		} else {
			minusOne = append(minusOne, i)
		}
	}
	// Order Normal Number
	for i := 0; i < len(normalNumber)-1; {
		if normalNumber[i] > normalNumber[i+1] {
			normalNumber[i], normalNumber[i+1] = normalNumber[i+1], normalNumber[i]
			i = 0
		} else {
			i++
		}
	}
	// Insert -1 at index
	for i := 0; i < len(minusOne); i++ {
		normalNumber = append(normalNumber, -1)
		if minusOne[i] == 0 {
			copy(normalNumber[minusOne[i]+1:], normalNumber[:])
		} else {
			copy(normalNumber[minusOne[i]:], normalNumber[minusOne[i]-1:])
		}
		normalNumber[minusOne[i]] = -1
	}
	return normalNumber
}
