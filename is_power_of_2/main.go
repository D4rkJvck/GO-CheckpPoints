package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		num, err := strconv.Atoi(os.Args[1])
		if err == nil {
			fmt.Println(isPowerOf2(num))
		}
	}
}
