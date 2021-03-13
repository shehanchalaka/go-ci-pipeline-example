package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2, 3) != 5 {
		t.Fatal("Someone has goofed")
	}
}

func TestSquare(t *testing.T) {
	if Square(5) != 25 {
		t.Fatal("Earth is square")
	}
}
