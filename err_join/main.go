package main

import (
	"cmp"
	"strconv"
)

func sum(a, b string) (int, error) {
	x, e1 := strconv.Atoi(a)
	y, e2 := strconv.Atoi(b)

	if err := cmp.Or(e1, e2); err != nil {
		return 0, err
	}
	return x + y, nil
}

func main() {
	s, _ := sum("1", "2")
	println("1+2", "=", s)
}
