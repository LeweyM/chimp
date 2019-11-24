package Repl

import (
	"Chimp/Evaluator"
	"Chimp/Lexer"
	"Chimp/Object"
	"Chimp/Parser"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := Object.NewEnvironment(nil)

	for {
		color.Blue("Go on...")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		text := scanner.Text()
		lexer := Lexer.New(text)
		parser := Parser.New(*lexer)

		programme := parser.ParseProgramme()

		if errors := parser.GetErrors(); len(errors) > 0 {
			color.Set(color.FgRed)
			_, _ = io.WriteString(out, "Parsing Error:\n")
			for i, err := range errors {
				_, _ = io.WriteString(out, fmt.Sprintf("%d: %s\n", i, err))
			}
			color.Unset()
		} else {
			p, err := Evaluator.Eval(programme, env)

			if err != nil {
				color.Set(color.FgRed)
				_, _ = io.WriteString(out, "Evaluator error:\n")
				_, _ = io.WriteString(out, err.Error())
				_, _ = io.WriteString(out, "\n")
				color.Unset()
			} else {
				color.Set(color.FgYellow)
				_, _ = io.WriteString(out, "$: ")
				_, _ = io.WriteString(out, p.Inspect())
				_, _ = io.WriteString(out, "\n")
				color.Unset()
			}
		}
	}
}
