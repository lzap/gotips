package main

import "reflect"

// Platform: Go on Intel x86_64
func main() {
	s := struct{}{}
	size := reflect.TypeOf(s).Size()
	println("Size is:", size, "bytes")
}
