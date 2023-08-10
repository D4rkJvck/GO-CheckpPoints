package main

func contain(str, l string) bool {
	for i := range str {
		if string(str[i]) == l {
			return true
		}
	}
	return false
}
