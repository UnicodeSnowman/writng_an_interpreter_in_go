package main

import (
	"fmt"
	"os"

	"github.com/unicodesnowman/writing_an_interpreter_in_go/repl"
)

func main() {
	fmt.Printf("Hello! Welcome! REPL!\n")
	repl.Start(os.Stdin, os.Stdout)
}
