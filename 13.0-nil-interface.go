package main

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
func CallMeMaybe(n Namer, name string) {
	if n == nil {
		return
	}
	n.SetName(name)
}

func main() {
	var namer Namer //= &People{firstName: "Pa", lastName: "In"}
	CallMeMaybe(namer, "letme")
}

// END OMIT
