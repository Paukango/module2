package main

import (
	"fmt"
)

func main() {
	var R *float64
	var pi float64 = 3.14
	var C float64 = 35
	R = &C
	fmt.Print(*R, "\n")
	*R = *R / 2 * pi
	fmt.Print(*R)
}
