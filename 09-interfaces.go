package main

// START OMIT
type Namer interface {
	Name() string
	//SetName(string) string
	SetName(newName string) (name string)
}
type People struct {
	firstName string
	lastName  string
}

func (a People) Name() string {
	return a.firstName + " " + a.lastName
}
func (a People) SetName(string) string {
	return "string"
}

// END OMIT
func main() {

}
