package main

type Person struct {
	Name string
	Age  int
}

func fByValue(p Person) {}

func fByReference(p *Person) {}

func main() {
	p := Person{"John", 30}

	fByValue(p)
	fByReference(&p)
}
