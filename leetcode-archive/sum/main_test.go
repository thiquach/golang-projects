package main

import "testing"

// Testing Sum function in main.go
// main_test.go contains tests for code in main.go
// functions in main.go to be capitalised and exported for visibility in main_test.go
func TestSum(t *testing.T) {
	result := Sum(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}
