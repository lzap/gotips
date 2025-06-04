package main

var unexported string

var Exported string

func Unexported() string {
	return unexported
}

func SetUnexported(value string) {
	unexported = value
}
