package main

import (
	"fmt"
	"reflect"
)

// START OMIT
type Super interface {
	Sky() string
}

type Not struct{}

func (n Not) Sky() string {
	return "wrath"
}
func main() {
	var n = Not{}
	whichType(n)
}

func whichType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case Super:
		fmt.Println(v.Sky())
	default:
		fmt.Println(reflect.TypeOf(i).String())
	}
}

// END OMIT
