package main

import "errors"

func Process() {
	err := Do()

	if errors.Is(err, ErrSomething) {
		println("Error value", err.Error())
	}

	var sErr *SomethingError
	if errors.As(err, &sErr) {
		println("Error type", err.Error())
	}
}
