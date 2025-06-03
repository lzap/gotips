package main

import "cmp"

func main() {
	a := 42
	b := cmp.Or(a, 10)
	println(a, b)
}
