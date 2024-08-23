package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Tyler-Meador/monkey-lang-go/monkey-lang-go/lexer"
	"github.com/Tyler-Meador/monkey-lang-go/monkey-lang-go/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
