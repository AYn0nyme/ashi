package main

func main() {
	Lexer("def string stupid = 5+5")
}

// c'est pas efficace ce que je fais, si je fais mot par mot, je vais miss les 5+5
// donc faut faire char par char...
