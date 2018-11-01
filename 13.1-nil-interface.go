package main

import (
	"errors"
	"fmt"
)

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

// START OMIT
func CallMeMaybe(n Namer, name string) error {
	if n == nil {
		return errors.New("'n' cannot be nil")
	}
	n.SetName(name)
	return nil
}

func main() {
	var namer Namer //= &People{firstName: "Pa", lastName: "In"}
	err := CallMeMaybe(namer, "letme")
	if err != nil {
		fmt.Println(err)
		// some actions
		return
	}
	fmt.Println("I am crazy")
}

// END OMIT
