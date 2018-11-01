package main

import (
	"fmt"
)

// START OMIT
type Namer interface {
	SetName(string) string
}
type People struct {
	firstName string
	lastName  string
}

func (a *People) SetName(name string) string {
	a.firstName = name
	return a.firstName + " " + a.lastName
}
func CallMeMaybe(n Namer, name string) {
	n.SetName(name)
}

func main() {
	var namer Namer //= &People{firstName: "Pa", lastName: "In"}
	CallMeMaybe(namer, "letme")
	fmt.Printf("Value: %v; Type: %T;\n", namer, namer)
}

// END OMIT
