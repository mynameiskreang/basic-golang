package main

import "fmt"

//* at least 6 charactors long
//* increase & decrese number is not allowed (at lease 3 chars)
//ex: 123456, 345678, 765432, 9876543
//* repeat number (at lease 3 chars)
//ex: 111111, 5678880, 1923339
//* repeat 2 digits but sequence in set
//ex: 1112233, 445566, 887766
func main() {
	//input := []string{"879123456879", "345678", "765432", "9876543"}
	//input := []string{"111333", "222333", "333444", "444555"}
	input := []string{"879112233879", "445566", "887766"}
	//input := []string{"152535", "415161", "817161"}
	isPass := case1(input)
	fmt.Println(isPass)
	isPass = case2(input)
	fmt.Println(isPass)
	isPass = case3(input)
	fmt.Println(isPass)
}

func case1(input []string) bool {
	// case 1:
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if j-2 > 0 && j+2 < len(input[i]) {
				iBy := int(input[i][j])
				fp := int(input[i][j+1])
				sp := int(input[i][j+2])
				if iBy+1 == fp && iBy+2 == sp {
					return false
				}
				fm := int(input[i][j-1])
				sm := int(input[i][j-2])
				if iBy-1 == fm && iBy-2 == sm {
					return false
				}
			}
		}
	}
	return true
}

func case2(input []string) bool {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if j-2 > 0 && j+2 < len(input[i]) {
				iBy := input[i][j]
				fp := input[i][j+1]
				sp := input[i][j+2]
				if iBy == fp && iBy == sp {
					return false
				}
				fm := input[i][j-1]
				sm := input[i][j-2]
				if iBy == fm && iBy == sm {
					return false
				}
			}
		}
	}
	return true
}

func case3(input []string) bool {
	flag := 0
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for c := 0; c < len(numbers); c++ {
		cNumber := numbers[c]
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				//if j-1 > 0 && j+1 < len(input[i]) {
				if j+1 < len(input[i]) {
					cBy := input[i][j]
					nCBy := input[i][j+1]
					if string(cBy) == cNumber && string(nCBy) == cNumber {
						flag++
						break
					}
				}
			}
			if flag > 3 {
				return false
			}
		}
	}
	return true
}
