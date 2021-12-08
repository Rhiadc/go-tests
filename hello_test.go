package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Person")
	want := "Hello, Person"

	if result != want {
		t.Errorf("Result: '%s, Waiting: '%s'", result, want)
	}
}
