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
