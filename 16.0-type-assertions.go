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

	d := isThisPython[0].(int)
	o := isThisPython[1].(string)
	t := isThisPython[2].(bool)
	a, ok := isThisPython[2].(int)
	if !ok {
		fmt.Println("not today")
	}

	fmt.Println(d, o, t, a)
}

// END OMIT
