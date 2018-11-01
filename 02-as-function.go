package main

import "fmt"

// START OMIT
type Pipl struct {
	firstName string
	lastName  string
}

func Name(a Pipl) string {
	return a.firstName + " " + a.lastName
}

func main() {
	var a = Pipl{
		firstName: "codename",
		lastName:  "47",
	}
	fmt.Println(Name(a))
}

// END OMIT
