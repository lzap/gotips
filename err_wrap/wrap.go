package main

import (
	"errors"
	"fmt"
)

var ErrSomething = errors.New("something went wrong")

type SomethingError string
func (e SomethingError) Error() string {return string(e)}

func Do() error {
	return fmt.Errorf("do failed: %w", ErrSomething)
}
