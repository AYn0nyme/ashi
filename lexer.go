package main

import (
	"fmt"
	"strings"
)

var Tokens []Token
var keywords = []string{
	"def",
	"return",
	"continue",
	"until",
	"take",
	"give",
	"for",
	"switch",
	"fun",
	"from",
	"not",
	"case",
	"default",
	"do",
}
var typesList = []string{
	"string",
	"number",
	"float",
	"boolean",
	"null",
	"char",
}
var operators = []string{
	"*",
	"/",
	"-",
	"+",
	"%",
}

func Lexer(SourceCode string) {

	lines := strings.Lines(SourceCode)

	for line := range lines {
		words := strings.Fields(line)

		for _, word := range words {
			if Has(keywords, word) {
				Tokens = append(Tokens, Token{Type: "keyword", Value: word})
			} else if Has(typesList, word) {
				Tokens = append(Tokens, Token{Type: "type", Value: word})
			} else if Has(operators, word) {
				Tokens = append(Tokens, Token{Type: "operator", Value: word})
			}

		}
	}

	fmt.Println(Tokens)
}
