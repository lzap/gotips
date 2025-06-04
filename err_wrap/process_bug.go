package main

import "errors"

var ErrQueueFull = errors.New("queue is full")

func main() {
	// START OMIT
	err := DoSomething()

	if errors.Is(err, ErrSomething) {
		println("Error value", err.Error())
	}

	var sErr SomethingError
	if errors.As(err, &sErr) {
		println("Error type", err.Error())
	}
	// END OMIT
}
