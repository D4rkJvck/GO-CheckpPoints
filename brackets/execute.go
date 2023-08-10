package main

import (
	"fmt"
)

func Execute(s, a string) {
	fmt.Println("Sample?")
	fmt.Scan(&s)
	fmt.Println(check(s))
	// Reload
	fmt.Println("\nAnother try? y or n")
	fmt.Scan(&a)
	switch a {
	case "y", "Y":
		Execute(s, a)
	case "n", "N":
		return
	default:
		fmt.Println("Wrong answer!")
		return
	}

}
