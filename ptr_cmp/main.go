package main

import "cmp"

func main() {
	a := 42
	b := cmp.Or(0, 0, 10)
	println(a, b)
}
