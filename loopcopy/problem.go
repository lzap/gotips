package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{{"Alice", 30}, {"Bob", 25}}

	for _, p := range people {
		p.Age++ // HLcopy
	}

	fmt.Println(people)
}
