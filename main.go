package main

import (
	"fmt"
	"lib/lexer"
)

func main() {
	fmt.Println("Hello world")
	input := "let abc = 1;"
	l := lexer.New(input)
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
}
