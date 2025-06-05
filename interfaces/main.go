package main

import "testing"

type Closer interface {
	Close()
}

type resource struct{}

func (r *resource) Close() {
	// do something
}

// ensure that resource implement the "Closer" interface
var _ = Closer(&resource{})

// fakeResource can now be used in testing
type fakeResource struct {
	closeCalls int
}

func (f *fakeResource) Close() {
	f.closeCalls++
}

func fooThatShouldCloseTheResource(c Closer) {
	// do something that eventually needs close c
}
func TestFooCloses(t *testing.T) {
	var fr fakeResource

	fooThatShouldCloseTheResource(&fr)
	// ensure the resource is closed exactly once
	if fr.closeCalls != 1 {
		t.Fatalf("close should be called exactly once, got %v", fr.closeCalls)
	}
}
