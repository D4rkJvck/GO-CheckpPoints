package main

func GetGCD(x, y int) int {
	if x%y == 0 {
		return y
	}
	return GetGCD(y, x%y)
}
