package Repl

import (
	"Chimp/Evaluator"
	"Chimp/Lexer"
	"Chimp/Object"
	"Chimp/Parser"
	"bufio"
	"fmt"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := Object.NewEnvironment()

	for {
		fmt.Println("Go on...")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		text := scanner.Text()
		lexer := Lexer.New(text)
		parser := Parser.New(*lexer)

		programme := parser.ParseProgramme()

		e := Evaluator.Eval(programme, *env)

		io.WriteString(out, "Evaluated code: ")
		io.WriteString(out, e.Inspect())
		io.WriteString(out, "\n")

	}
}
