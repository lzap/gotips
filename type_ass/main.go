package main

func main() {
	i := 42
	var x any = "any"

	a := string(i)
	b, ok := x.(string)

	println(a, b, ok)
}
