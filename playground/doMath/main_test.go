package main

import "testing"

// main_test.go contains tests for code in main.go
// functions in main.go to be capitalised and exported for visibility in main_test.go
func TestDoMath(t *testing.T) {
	result := DoMath(2, 2, add)
	expected := 4

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}

	result = DoMath(46, 5, subtract)
	expected = 41

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}
