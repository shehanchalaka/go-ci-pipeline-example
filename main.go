package main

import (
	"fmt"

	"github.com/shehanchalaka/go-ci-pipeline-example/book"
)

func main() {
	fmt.Println("CI/CD example")

	msg := book.Book()
	fmt.Println(msg)
}
