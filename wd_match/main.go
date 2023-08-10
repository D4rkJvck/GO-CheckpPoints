package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		if match(os.Args[1], os.Args[2]) {
			fmt.Println(os.Args[1])
		}
	}
}
