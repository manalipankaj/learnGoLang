package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello World"
	actual := helloWorld()

	if actual != expected {
		t.Error("test failed")
	}
}