package main

import "errors"

func main() {
	myString := "hello"
	myNil := nil // HLnil
	myErr := errors.New("e1")
	myInt, myErr := 2, errors.New("e2")
	myErr := errors.New("e3") // HLerr

	println(myString, myInt, myErr)
}
