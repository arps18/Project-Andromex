package main

import (
	"fmt" // For formatted I/O ops
	"os"  // For OS interface (accessing CLI args)
)

func main(){
	fmt.Println("Andromex REPL(ctrl + c to exit)")
	_ = os.Args
}