package main

import (
	"fmt"
)

// START OMIT

func main() {
	var isThisPython []interface{}
	var n int
	var s string
	var b bool
	isThisPython = append(isThisPython, n)
	isThisPython = append(isThisPython, s)
	isThisPython = append(isThisPython, b)
	fmt.Println(isThisPython)
}

// END OMIT
