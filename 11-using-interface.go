package main

import (
	"fmt"
)

// START OMIT
type Namer interface {
	Name() string
}
type People struct {
	firstName string
	lastName  string
}

func (a People) Name() string {
	return a.firstName + " " + a.lastName
}

func CallMeMaybe(n Namer) {
	fmt.Printf("Value: %v; Type: %T;\n", n, n)
}

func main() {
	var namer Namer = &People{firstName: "Pa", lastName: "In"}
	CallMeMaybe(namer)
	var people = People{firstName: "Kil", lastName: "Ler"}
	CallMeMaybe(people)
}

// END OMIT
