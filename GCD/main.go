package main

import (
	"fmt"
)

func main() {
	if len(os.Args) == 3 {
		//First Input
		fmt.Println("First number?")
		var first int
		fmt.Scan(&first)
		//Second Input
		fmt.Println("Second number?")
		var second int
		fmt.Scan(&second)
		//Final Result
		fmt.Printf("Greatest Common Divider: %d", GetGCD(first, second))
	} else {
		fmt.Println("Minimum 2 Numbers")
	}
}
