package main

func From[T any](ref *T) (value T) {
	if ref != nil {
		value = *ref
	}
	return
}

func main() {
	i := 42
	ptr := &i
	// _ = *ptr
	i = From(ptr)
	println(i)
}
