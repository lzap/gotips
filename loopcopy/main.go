package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{{"Alice", 30}, {"Bob", 25}}

	for i := range people {
		people[i].Age++
	}
	
	fmt.Println(people)
}
