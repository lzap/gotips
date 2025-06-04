package main

func main() {
	var s []string // = nil

	println(len(s))

	s = append(s, "hello")

	println(s)
}
