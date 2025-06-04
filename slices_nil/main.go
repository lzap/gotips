package main

func main() {
	var s []string // = nil

	println(len(s))

	for i, v := range s {
		println(i, v)
	}

	s = append(s, "hello")

	println(s)
}
