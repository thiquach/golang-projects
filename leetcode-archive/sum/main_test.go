package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(2, 5)
	expected := 7

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}
