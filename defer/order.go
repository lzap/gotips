package main

func main() {
	i := 1
	defer println(i)
	i++	
	defer println(i)
	i = 42
}
