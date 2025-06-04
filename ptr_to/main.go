package main

func To[T any](value T) *T {
	return &value
}

func main() {
	// ptr := &1
	ptr := To(1)
	println(*ptr)
}
