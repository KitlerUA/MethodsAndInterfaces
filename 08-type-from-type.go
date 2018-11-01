package main

import (
	"fmt"
)

// START OMIT
type Pipl int

func (a *Pipl) SeePlusPlus() {
	*a += Pipl(1)
}

func main() {
	var a Pipl = 665
	a.SeePlusPlus()
	fmt.Println(a)
}

// END OMIT
