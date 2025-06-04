package main

func FromOr[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}

func Or[T any](value *T, defaultValue T) *T {
	if value == nil {
		return &defaultValue
	}
	return value
}

func main() {
	b := true
	ptr := &b
	value := FromOr(ptr, false)
	println(value)
}
