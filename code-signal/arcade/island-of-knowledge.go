package main

// 19: areEquallyStrong https://app.codesignal.com/arcade/intro/level-5/g6dc9KJyxmFjB98dL
func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	if yourLeft+yourRight != friendsLeft+friendsRight {
		return false
	}

	mapStrong := make(map[int]bool)
	mapStrong[yourLeft] = true
	mapStrong[yourRight] = true
	mapStrong[friendsLeft] = true
	mapStrong[friendsRight] = true
	if len(mapStrong) > 2 {
		return false
	}

	return true
}
