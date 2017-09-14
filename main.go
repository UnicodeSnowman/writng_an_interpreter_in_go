package main

import (
	"fmt"
	"os"

	"github.com/unicodesnowman/monkey/repl"
)

func main() {
	fmt.Printf("Hello! Welcome! REPL!\n")
	repl.Start(os.Stdin, os.Stdout)
}
