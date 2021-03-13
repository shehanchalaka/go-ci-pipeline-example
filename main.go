package main

import (
	"fmt"

	"github.com/shehanchalaka/go-ci-pipeline-example/book"
)

func Calculate(x int, y int) int {
	return x + y
}

func Square(x int) int {
	return x * x
}

func main() {
	fmt.Println("Go CI/CD example")

	msg := book.Book()
	fmt.Println(msg)
}
