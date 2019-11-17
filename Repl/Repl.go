package Repl

import (
	"Chimp/Lexer"
	"Chimp/Parser"
	"bufio"
	"fmt"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		io.WriteString(out, "Parsed code: ")
		io.WriteString(out, programme.ToString())
		io.WriteString(out, "\n")

	}
}
