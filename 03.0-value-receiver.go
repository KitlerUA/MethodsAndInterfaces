package main

import "fmt"

// START OMIT
type Pipl struct {
	firstName string
	lastName  string
}

func (a Pipl) Name() string {
	return a.firstName + " " + a.lastName
}

func (a Pipl) SetFirstName(first string) {
	a.firstName = first
}

func main() {
	var a = &Pipl{
		firstName: "codename",
		lastName:  "47",
	}
	a.SetFirstName("rosswell")
	fmt.Println(a.Name())
	(*a).SetFirstName("19")
	fmt.Println(a.Name())
}

// END OMIT
