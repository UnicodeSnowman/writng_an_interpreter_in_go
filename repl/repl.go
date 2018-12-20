package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/unicodesnowman/writing_an_interpreter_in_go/evaluator"
	"github.com/unicodesnowman/writing_an_interpreter_in_go/lexer"
	"github.com/unicodesnowman/writing_an_interpreter_in_go/object"
	"github.com/unicodesnowman/writing_an_interpreter_in_go/parser"
)

const PROMPT = ">> "

func goodbye(out io.Writer) {
	fmt.Fprintf(out, "Goodbye!\n")
}

func Start(in io.Reader, out io.Writer) {
	defer goodbye(out)
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)

		// TODO support up/down arrow for scrolling through history
		// handle automatically via rlwrap?
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			goodbye(out)
			os.Exit(0)
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
