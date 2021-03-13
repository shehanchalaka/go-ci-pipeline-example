package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2, 3) != 5 {
		t.Fatal("Someone has goofed")
	}
}
