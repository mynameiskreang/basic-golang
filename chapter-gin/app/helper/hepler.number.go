package helper

import (
	"strconv"
)

func StringToInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		// return default value
		return 0
	}
	return output
}
