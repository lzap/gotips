package main

func To[T any](value T) *T {
	return &value
}

func ToErr[T any](value T, e error) (*T, error) {
	return &value, e
}

func main() {
	// ptr := &1
	ptr := To(1)
	println(*ptr)
}
