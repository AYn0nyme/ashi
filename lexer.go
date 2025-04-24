package main

import (
	"slices"
	"strings"
	"unicode"
)

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
var tmp string
var operators = []string{
	"*",
	"/",
	"-",
	"+",
	"%",
	"=",
}

func Lexer(SourceCode string) []Token {
	var Tokens []Token

	lines := strings.Lines(SourceCode)

	for line := range lines {
		for _, char := range line {
			if unicode.IsLetter(char) {
				if GetType(tmp) != "identifier" {
					Tokens = append(Tokens, Token{
						Type:  GetType(tmp),
						Value: tmp,
					})
					tmp = ""
				}
				tmp += string(char)
			} else if unicode.IsNumber(char) {
				tmp = ""
				Tokens = append(Tokens, Token{
					Type:  "number",
					Value: int(char - '0'),
				})
			} else if unicode.IsSpace(char) {
				Tokens = append(Tokens, Token{
					Type:  GetType(tmp),
					Value: tmp,
				})
				tmp = ""
			} else {
				if GetType(tmp) != "identifier" && GetType(string(char)) != "operator" {
					Tokens = append(Tokens, Token{
						Type:  GetType(tmp),
						Value: tmp,
					})
					tmp = ""
				}
				tmp += string(char)
			}
		}
	}
	return Tokens
}

func GetType(str string) string {
	if slices.Contains(keywords, str) {
		return "keyword"
	}
	if slices.Contains(typesList, str) {
		return "type"
	}
	if slices.Contains(operators, str) {
		return "operator"
	}
	return "identifier"
}

func NaN(str string) bool {
	for _, v := range str {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}
