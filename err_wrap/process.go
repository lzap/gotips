package main

import "errors"

func Process() {
	// START OMIT
	err := DoSomething()

	if errors.Is(err, ErrSomething) {
		println("Error value", err.Error())
	}

	var sErr *SomethingError // HLref
	if errors.As(err, &sErr) {
		println("Error type", err.Error())
	}
	// END OMIT
}
