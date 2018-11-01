package main

// START OMIT
type Pipl struct {
	firstName string
	lastName  string
}

func (a Pipl) Name() string {
	return a.firstName + " " + a.lastName
}

func takeMeEverything(interface{}) {
	return
}

func main() {
	var p Pipl
	takeMeEverything(p)
}

// END OMIT
