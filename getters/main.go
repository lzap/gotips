package main

var unexported string

var Exported string

func do() {
	println("Done")
}

func Unexported() string {
	return unexported
}

func SetUnexported(value string) {
	unexported = value
}

func main() {}
