package main

func f(i int) {
	if i--; i <= 0 {
		return
	}
	f(i)
}

func main() {
	var val int

	println(&val)
	f(1000)
	println(&val)
}
