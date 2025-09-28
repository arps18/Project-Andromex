package main

import (
	"fmt" // For formatted I/O ops
	"os"  // For OS interface (accessing CLI args)

	"github.com/arps18/andromex/lexer"
)

func main() {
	fmt.Println("Andromex REPL(ctrl + c to exit)")
	_ = os.Args

	// Test lexer
	input := "2 + 3 * 4"
	l := lexer.New(input)
	for tok := l.NextToken(); tok.Type != "EOF"; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
