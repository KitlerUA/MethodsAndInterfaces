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

func SetFirstName(a *Pipl, first string) {
	a.firstName = first
}

func main() {
	var a = Pipl{
		firstName: "codename",
		lastName:  "47",
	}
	SetFirstName(&a, "rosswell")
	fmt.Println(a.Name())
}

// END OMIT
