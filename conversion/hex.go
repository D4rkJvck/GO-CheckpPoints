package main

import "fmt"

func main() {
	fmt.Println(Hex(255))
}

func Hex(n int) string {
	if n == 0 {
		return 
	}
	t := "0123456789abcdef"
	return Hex(n/16) + string(t[n%16])
}

func Palindromer(s string) string {
	for _, letter := range s {

	}
}

func HasPal(s string) bool {
	
}