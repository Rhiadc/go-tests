package main

import "testing"

func TestHello(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, want string) {
		t.Helper()
		if result != want {
			t.Errorf("Result: '%s, Waiting: '%s'", result, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Person")
		want := "Hello, Person"
		verifyCorrectMessage(t, result, want)
	})

	t.Run("say 'Hello, Go' when an empty string is passed", func(t *testing.T) {
		result := Hello("")
		want := "Hello, Go"
		verifyCorrectMessage(t, result, want)
	})
}
