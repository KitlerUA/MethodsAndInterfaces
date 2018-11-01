package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var isThisPython []interface{}
	var n int
	var s string
	var b bool
	isThisPython = append(isThisPython, n)
	isThisPython = append(isThisPython, s)
	isThisPython = append(isThisPython, b)

	// START OMIT
	for i := range isThisPython {
		fmt.Printf("%s ", toString(isThisPython[i]))
	}
}

func toString(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	// case string:
	// 	return v
	default:
		return reflect.TypeOf(i).String()
	}
}

// END OMIT
