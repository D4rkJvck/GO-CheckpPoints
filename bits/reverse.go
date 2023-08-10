package main

import (
	"fmt"
)

func main() {
	var oct, byt byte
	fmt.Println("Byte?")
	fmt.Scan(&oct)
	fmt.Printf("%b\n%b", oct, BitReverse(oct, byt, 8))
}

func BitReverse(oct, byt byte, num int) byte {
	if num == 0 {
		return byt
	}
	byt = byt*2 + oct%2
	oct /= 2
	return BitReverse(oct, byt, num-1)
}
