package main

import (
	"fmt"
	"strconv"
	"unicode"
)

//48: isDigit https://app.codesignal.com/arcade/intro/level-11/Zr2XXEpkBPtYWqPJu
func isDigit(symbol string) bool {
	return unicode.IsDigit(rune(symbol[0]))
}

//49: lineEncoding https://app.codesignal.com/arcade/intro/level-11/o2uq6eTuvk7atGadR
func lineEncoding(s string) string {
	checker := rune(s[0])
	result := ""
	count := 0
	for _, r := range s {
		if checker == r {
			count++
			continue
		} else {
			if count > 1 {
				result = result + fmt.Sprintf("%d%c", count, checker)
			} else {
				result = result + fmt.Sprintf("%c", checker)
			}
			count = 1
			checker = r
		}
	}
	if count > 1 {
		result = result + fmt.Sprintf("%d%c", count, checker)
	} else {
		result = result + fmt.Sprintf("%c", checker)
	}
	return result
}

// 50: chessKnight https://app.codesignal.com/arcade/intro/level-11/pwRLrkrNpnsbgMndb
func chessKnight(cell string) int {
	mapChar := make(map[byte]int, 8)
	mapChar['a'] = 1
	mapChar['b'] = 2
	mapChar['c'] = 3
	mapChar['d'] = 4
	mapChar['e'] = 5
	mapChar['f'] = 6
	mapChar['g'] = 7
	mapChar['h'] = 8
	cellI, _ := mapChar[cell[0]]
	cellJ, _ := strconv.Atoi(string(cell[1]))
	dataSetI := []int{2, 1, -1, -2, -2, -1, 1, 2}
	dataSetJ := []int{1, 2, 2, 1, -1, -2, -2, -1}
	count := 0
	for s := 0; s < len(dataSetI); s++ {
		if cellI+dataSetI[s] > 0 && cellJ+dataSetJ[s] > 0 && cellI+dataSetI[s] <= 8 && cellJ+dataSetJ[s] <= 8 {
			count++
		}
	}
	return count
}

// 51: deleteDigit https://app.codesignal.com/arcade/intro/level-11/vpfeqDwGZSzYNm2uX
func deleteDigit(n int) int {
	max := 0
	nS := strconv.Itoa(n)
	for i := 1; i < len(nS)+1; i++ {
		newNS := nS[:i-1] + nS[i:]
		nSI, _ := strconv.Atoi(newNS)
		if nSI > max {
			max = nSI
		}
	}
	return max
}
