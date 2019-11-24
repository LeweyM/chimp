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
	env := Object.NewEnvironment(nil)

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

		if errors := parser.GetErrors(); len(errors) > 0 {
			io.WriteString(out, "Parsing Error:\n")
			for i, err := range errors {
				io.WriteString(out, fmt.Sprintf("%d: %s\n", i, err))
			}
		} else {
			p, err := Evaluator.Eval(programme, env)

			if err != nil {
				io.WriteString(out, "Evaluator error:\n")
				io.WriteString(out, err.Error())
				io.WriteString(out, "\n")
			} else {
				io.WriteString(out, "Evaluated code: ")
				io.WriteString(out, p.Inspect())
				io.WriteString(out, "\n")
			}

		}
	}
}
