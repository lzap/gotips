package main

import (
	"errors"
	"fmt"
)

// START OMIT
var ErrSomething = errors.New("something went wrong")

type SomethingError string

func (e SomethingError) Error() string { return string(e) }

func DoSomething() error {
	return fmt.Errorf("do failed: %w", ErrSomething)
}

// END OMIT
