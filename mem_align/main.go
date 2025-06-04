package main

import "reflect"

// Platform: Go on Intel x86_64
func main() {
	s := struct {
		a int64 // 8 bytes
		b bool  // 1 byte
	}{}
	size := reflect.TypeOf(s).Size()
	println("Size is:", size, "bytes")
}
