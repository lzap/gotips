package main

// START OMIT
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
	b := false
	println(FromOr(&b, true))
}

// END OMIT
