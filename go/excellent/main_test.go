package main

import "testing"

func TestEven0r0dd(t *testing.T) {
	result := Even0r0dd(10)
	if result != "Even" {
		t.Errorf("Expected Even, actual: %s", result)
	}
}