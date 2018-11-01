package main

type Pipl struct {
	firstName string
	lastName  string
}
type People struct {
	firstName string
	lastName  string
}

// START OMIT

func (a Pipl) Name() string {
	return a.firstName + " " + a.lastName
}

func (a People) Name() string {
	return a.firstName + " " + a.lastName
}

// END OMIT

func main() {

}
