package main

func isPowerOf2(n int) string {
	if n < 2 {
		return "true"
	}
	if n%2 != 0 {
		return "false"
	}
	return isPowerOf2(n / 2)
}
