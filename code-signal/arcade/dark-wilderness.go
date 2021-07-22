package main

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
