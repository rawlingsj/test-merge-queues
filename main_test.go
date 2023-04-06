package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	actual := HelloWorld()

	if expected != actual {
		t.Errorf("Expected: %s, got: %s", expected, actual)
	}
}
