package main

func f() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func main() {
	println(f())
}
