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
	ptr := To(true)
	value := FromOr(ptr, false)
	println(value)
}
