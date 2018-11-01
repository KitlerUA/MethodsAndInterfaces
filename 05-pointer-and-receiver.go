package main

type Pipl struct {
	firstName string
	lastName  string
}

func (a Pipl) Name() string {
	return a.firstName + " " + a.lastName
}

// START OMIT

func (a Pipl) SetFirstName(first string) {
	a.firstName = first
}

func (a *Pipl) SetFirstName(first string) {
	a.firstName = first
}

// END OMIT

func main() {

}
