package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/unicodesnowman/monkey/lexer"
	"github.com/unicodesnowman/monkey/token"
)

const PROMPT = ">> "

func goodbye(out io.Writer) {
	fmt.Fprintf(out, "Goodbye!\n")
}

func Start(in io.Reader, out io.Writer) {
	defer goodbye(out)
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
