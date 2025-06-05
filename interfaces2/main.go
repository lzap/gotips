package main

import "testing"

type Writer interface {
	Write()
}

type Closer interface {
	Close()
}

// small interfaces can be combined
type WriteCloser interface {
	Writer
	Closer
}

