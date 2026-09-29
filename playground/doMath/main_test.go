package main

import "testing"

// TestDoMath() in upper camel casing
func TestDoMath(t *testing.T) {
	got := doMath(2, 2, add)
	expected := 4

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}

	got = doMath(46, 5, subtract)
	expected = 41

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}
