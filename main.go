package main

import "fmt"

func main() {
	tokens := Lexer("def number index = index++")
	fmt.Println("def number index = index++", tokens)
}
