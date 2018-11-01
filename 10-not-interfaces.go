package main

// START OMIT
type Namer interface {
	Name() string
	//SetName(string) string
	SetName(newName string) (name string)
}

type Pipl struct {
	firstName string
	lastName  string
}

func (a Pipl) Name() string {
	return a.firstName + " " + a.lastName
}

// END OMIT

func main() {

}
