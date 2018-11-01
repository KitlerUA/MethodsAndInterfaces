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

func main() {
	var a = Pipl{
		firstName: "codename",
		lastName:  "47",
	}
	fmt.Println(a.Name())
}

// END OMIT
